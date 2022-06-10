package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
}

func TestWebhook(t *testing.T) {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:         scheme,
		LeaderElection: false,
	})
	assert.NoError(t, err)

	webhookServer := mgr.GetWebhookServer()
	webhookServer.Register("/validate-pod", &webhook.Admission{
		Handler: &ValidatingAdmission{},
	})

	ctx := ctrl.SetupSignalHandler()
	if err := mgr.Start(ctx); err != nil {
		assert.NoError(t, err)
	}
}
