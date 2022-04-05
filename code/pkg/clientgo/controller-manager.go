package clientgo

import (
	"context"
	"log"

	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type Controller struct {
	mgr ctrl.Manager
}

func (c *Controller) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
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

func (c *Controller) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, err error) {
	log.Println(req.NamespacedName)
	return ctrl.Result{}, nil
}

func StartController() error {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:         scheme,
		LeaderElection: false,
		//Namespace:          osArgs.WatchNamespace,
		//SyncPeriod:         &osArgs.SyncPeriod,
		//MetricsBindAddress: osArgs.MetricsAddr,
	})
	if err != nil {
		return err
	}

	if err := (&Controller{
		mgr: mgr,
	}).SetupWithManager(mgr, controller.Options{MaxConcurrentReconciles: 10}); err != nil {
		return err
	}

	// Setup the context that's going to be used in controllers and for the manager.
	ctx := ctrl.SetupSignalHandler()

	// +kubebuilder:scaffold:builder
	if err := mgr.Start(ctx); err != nil {
		return err
	}
	return nil
}
