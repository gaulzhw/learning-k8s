package shadow

import (
	"context"

	"github.com/go-logr/logr"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ reconcile.Reconciler = &ReconcileShadow{}

// +kubebuilder:rbac:groups=kingsport.k8s.io,resources=festivals,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kingsport.k8s.io,resources=festivals/status,verbs=get;update;patch

// ReconcileShadow reconciles a Shadow object
type ReconcileShadow struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

func (r *ReconcileShadow) SetupWithManager(mgr ctrl.Manager) error {
	return nil
}

func (r *ReconcileShadow) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	// TODO add logic
	return reconcile.Result{}, nil
}
