package clientgo

import (
	"context"
	"log"

	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	// +kubebuilder:scaffold:imports
)

func StartControllerWithIndex() error {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:         scheme,
		LeaderElection: false,
	})
	if err != nil {
		return err
	}

	mgr.GetFieldIndexer().IndexField(context.TODO(), &corev1.Pod{}, "spec.nodeName", func(obj client.Object) []string {
		pod, ok := obj.(*corev1.Pod)
		if !ok {
			return []string{}
		}

		if len(pod.Spec.NodeName) == 0 || pod.Status.Phase == corev1.PodSucceeded || pod.Status.Phase == corev1.PodFailed {
			return []string{}
		}

		return []string{pod.Spec.NodeName}
	})

	pods := &corev1.PodList{}
	err = mgr.GetAPIReader().List(context.TODO(), pods, client.MatchingFields{"spec.nodeName": "hub-control-plane"})
	if err != nil {
		return err
	}

	for _, pod := range pods.Items {
		log.Println(pod)
	}

	return nil
}
