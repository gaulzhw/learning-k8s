package informer

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wI2L/jsondiff"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// https://cloud.tencent.com/developer/article/1989055

var (
	scheme = runtime.NewScheme()
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

type PodController struct {
	client   client.Client
	recorder record.EventRecorder
}

func (c *PodController) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	// add index
	if err := mgr.GetFieldIndexer().IndexField(context.TODO(), &corev1.Pod{}, "labels", func(obj client.Object) []string {
		pod, ok := obj.(*corev1.Pod)
		if !ok {
			return []string{}
		}
		val := make([]string, 0, len(pod.Labels))
		for k, v := range pod.Labels {
			val = append(val, k+"/"+v)
		}
		return val
	}); err != nil {
		return err
	}

	c.recorder = mgr.GetEventRecorderFor("configmap-controller")

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
	if err := c.client.Get(ctx, req.NamespacedName, pod); err != nil {
		return ctrl.Result{}, err
	}
	//fmt.Printf("configmap: %v", pod)

	c.recorder.Event(pod, corev1.EventTypeNormal, "test-reason", "test-msg")
	return ctrl.Result{}, nil
}

func TestController(t *testing.T) {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:         scheme,
		LeaderElection: false,
	})
	assert.NoError(t, err)

	if err := (&PodController{
		client: mgr.GetClient(),
	}).SetupWithManager(mgr, controller.Options{
		MaxConcurrentReconciles: 1,
	}); err != nil {
		assert.NoError(t, err)
	}

	go func() {
		mgr.GetCache().WaitForCacheSync(context.TODO())
		pods := &corev1.PodList{}
		err := mgr.GetClient().List(context.TODO(), pods, client.MatchingFields{"labels": "component/etcd"})
		assert.NoError(t, err)
		for _, pod := range pods.Items {
			t.Log(pod.Namespace, pod.Name)
		}
	}()

	ctx := ctrl.SetupSignalHandler()
	if err := mgr.Start(ctx); err != nil {
		assert.NoError(t, err)
	}
}

func TestLabelCachedIndexController(t *testing.T) {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:         scheme,
		LeaderElection: false,
		NewCache:       newCache,
	})
	assert.NoError(t, err)

	if err := (&PodController{
		client: mgr.GetClient(),
	}).SetupWithManager(mgr, controller.Options{
		MaxConcurrentReconciles: 1,
	}); err != nil {
		assert.NoError(t, err)
	}

	// add index
	objs := []client.Object{
		&corev1.Pod{},
	}
	for _, obj := range objs {
		mgr.GetFieldIndexer().IndexField(context.TODO(), obj, indexName, func(obj client.Object) []string {
			val := make([]string, 0)
			for k, v := range obj.GetLabels() {
				val = append(val, indexKeyFun(k, v))
			}
			return val
		})
	}

	go func() {
		mgr.GetCache().WaitForCacheSync(context.TODO())
		pods := &corev1.PodList{}
		assert.NoError(t, mgr.GetClient().List(context.TODO(), pods, client.MatchingLabels{"component": "etcd"}))
		for _, pod := range pods.Items {
			t.Log(pod.Namespace, pod.Name)
		}
	}()

	ctx := ctrl.SetupSignalHandler()
	if err := mgr.Start(ctx); err != nil {
		assert.NoError(t, err)
	}
}

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

func TestForceDelete(t *testing.T) {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:         scheme,
		LeaderElection: false,
	})
	assert.NoError(t, err)

	ctx := context.TODO()
	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
		},
	}
	var period int64 = 0
	err = mgr.GetClient().Delete(ctx, cm, &client.DeleteOptions{
		GracePeriodSeconds: &period,
	})
	assert.NoError(t, err)
}
