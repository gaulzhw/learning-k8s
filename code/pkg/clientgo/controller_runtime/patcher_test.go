package controller_runtime

import (
	"context"
	"encoding/json"
	"testing"
	// +kubebuilder:scaffold:imports

	"github.com/stretchr/testify/assert"
	"github.com/wI2L/jsondiff"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func TestPatch(t *testing.T) {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:         scheme,
		LeaderElection: false,
		//Namespace:          osArgs.WatchNamespace,
		//SyncPeriod:         &osArgs.SyncPeriod,
		//MetricsBindAddress: osArgs.MetricsAddr,
	})
	assert.NoError(t, err)

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx",
			Namespace: "default",
		},
	}
	err = mgr.GetAPIReader().Get(context.TODO(), client.ObjectKeyFromObject(deploy), deploy)
	assert.NoError(t, err)

	newDeploy := deploy.DeepCopy()
	newDeploy.ObjectMeta.Labels["test1"] = "difftest1"

	oldObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(deploy)
	newObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(newDeploy)

	spec, _, err := unstructured.NestedFieldNoCopy(oldObj, "spec")
	t.Logf("spec: %v", spec)

	// 比较新旧deploy的不同，返回不同的bytes
	patch, err := jsondiff.Compare(oldObj, newObj)
	assert.NoError(t, err)

	// 打patch，patchBytes就是我们需要的了
	patchBytes, err := json.Marshal(patch)
	assert.NoError(t, err)

	err = mgr.GetClient().Patch(context.TODO(), deploy, client.RawPatch(types.JSONPatchType, patchBytes))
	assert.NoError(t, err)
}
