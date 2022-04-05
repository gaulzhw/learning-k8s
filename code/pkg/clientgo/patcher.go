package clientgo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/wI2L/jsondiff"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	// +kubebuilder:scaffold:imports
)

func Patch() error {
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

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx",
			Namespace: "default",
		},
	}
	err = mgr.GetAPIReader().Get(context.TODO(), client.ObjectKeyFromObject(deploy), deploy)
	if err != nil {
		return err
	}

	newDeploy := deploy.DeepCopy()
	newDeploy.ObjectMeta.Labels["test1"] = "difftest1"

	oldObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(deploy)
	newObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(newDeploy)

	spec, _, err := unstructured.NestedFieldNoCopy(oldObj, "spec")
	fmt.Printf("spec: %v", spec)

	// 比较新旧deploy的不同，返回不同的bytes
	patch, err := jsondiff.Compare(oldObj, newObj)
	if err != nil {
		return err
	}

	// 打patch，patchBytes就是我们需要的了
	patchBytes, err := json.Marshal(patch)
	if err != nil {
		return err
	}

	err = mgr.GetClient().Patch(context.TODO(), deploy, client.RawPatch(types.JSONPatchType, patchBytes))
	if err != nil {
		return err
	}

	return nil
}
