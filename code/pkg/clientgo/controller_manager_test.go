package clientgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
)

func TestController(t *testing.T) {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:         scheme,
		LeaderElection: false,
		NewCache: cache.BuilderWithOptions(cache.Options{
			DefaultSelector: cache.ObjectSelector{Field: fields.OneTermEqualSelector("metadata.namespace", "default")},
			SelectorsByObject: map[client.Object]cache.ObjectSelector{
				&corev1.Pod{}: {Field: fields.OneTermEqualSelector("metadata.namespace", "kube-system")},
			},
		}),
	})
	assert.NoError(t, err)

	if err := (&PodController{
		mgr: mgr,
	}).SetupWithManager(mgr, controller.Options{
		MaxConcurrentReconciles: 1,
	}); err != nil {
		assert.NoError(t, err)
	}

	if err := (&SecretController{
		mgr: mgr,
	}).SetupWithManager(mgr, controller.Options{
		MaxConcurrentReconciles: 1,
	}); err != nil {
		assert.NoError(t, err)
	}

	ctx := ctrl.SetupSignalHandler()
	if err := mgr.Start(ctx); err != nil {
		assert.NoError(t, err)
	}
}
