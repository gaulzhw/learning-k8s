package clientgo

import (
	"context"
	"log"

	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type PodController struct {
	mgr ctrl.Manager
}

func (c *PodController) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	if err := mgr.GetFieldIndexer().IndexField(context.TODO(), &corev1.Pod{}, "spec.nodeName", func(obj client.Object) []string {
		pod, ok := obj.(*corev1.Pod)
		if !ok {
			return []string{}
		}
		if len(pod.Spec.NodeName) == 0 || pod.Status.Phase == corev1.PodSucceeded || pod.Status.Phase == corev1.PodFailed {
			return []string{}
		}
		return []string{pod.Spec.NodeName}
	}); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(options).
		For(&corev1.Pod{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: func(e event.CreateEvent) bool {
				return true
			},
			UpdateFunc: func(e event.UpdateEvent) bool {
				return true
			},
			DeleteFunc: func(e event.DeleteEvent) bool {
				return true
			},
			GenericFunc: func(e event.GenericEvent) bool {
				return false
			},
		}).
		Complete(c)
}

func (c *PodController) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, err error) {
	pod := &corev1.Pod{}
	if err := c.mgr.GetClient().Get(ctx, req.NamespacedName, pod); err != nil {
		return ctrl.Result{}, err
	}
	log.Printf("pod: %v", pod)
	return ctrl.Result{}, nil
}

type SecretController struct {
	mgr ctrl.Manager
}

func (c *SecretController) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(options).
		For(&corev1.Secret{}).
		Complete(c)
}

func (c *SecretController) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, err error) {
	secret := &corev1.Secret{}
	if err := c.mgr.GetClient().Get(ctx, req.NamespacedName, secret); err != nil {
		return ctrl.Result{}, err
	}
	log.Printf("secret: %v", secret)
	return ctrl.Result{}, nil
}
