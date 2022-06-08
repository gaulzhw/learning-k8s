package informer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic/dynamicinformer"
	dynamicfake "k8s.io/client-go/dynamic/fake"
	"k8s.io/client-go/informers"
	clientsetfake "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"
)

func TestFakeInformer(t *testing.T) {
	ctx := context.Background()
	client := clientsetfake.NewSimpleClientset()

	// We will create an informer that writes added pods to a channel.
	pods := make(chan *corev1.Pod, 1)
	informers := informers.NewSharedInformerFactory(client, 0)
	podInformer := informers.Core().V1().Pods().Informer()
	podInformer.AddEventHandler(&cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*corev1.Pod)
			t.Logf("pod added: %s/%s", pod.Namespace, pod.Name)
			pods <- pod
		},
	})

	// Make sure informers are running.
	informers.Start(ctx.Done())
	cache.WaitForCacheSync(ctx.Done(), podInformer.HasSynced)

	// Inject an event into the fake client.
	p := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-pod",
		},
	}
	_, err := client.CoreV1().Pods("test-ns").Create(context.TODO(), p, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("error injecting pod add: %v", err)
	}

	select {
	case pod := <-pods:
		t.Logf("Got pod from channel: %s/%s", pod.Namespace, pod.Name)
		t.Logf("pod info: %+v", pod)
	case <-time.After(wait.ForeverTestTimeout):
		t.Error("Informer did not get the added pod")
	}
}

func TestFakeDynamicInformer(t *testing.T) {
	ctx := context.Background()
	scheme := runtime.NewScheme()
	corev1.AddToScheme(scheme)
	client := dynamicfake.NewSimpleDynamicClient(scheme)

	// prepare resource
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test",
		},
	}
	unstructuredPod, err := runtime.DefaultUnstructuredConverter.ToUnstructured(pod)
	assert.NoError(t, err)

	client.Resource(schema.GroupVersionResource{
		Group:    corev1.SchemeGroupVersion.Group,
		Version:  corev1.SchemeGroupVersion.Version,
		Resource: "pods",
	}).Namespace(pod.Namespace).Create(ctx, &unstructured.Unstructured{
		Object: unstructuredPod,
	}, metav1.CreateOptions{})

	informerFactory := dynamicinformer.NewDynamicSharedInformerFactory(client, time.Second*30)
	informerFactory.ForResource(schema.GroupVersionResource{
		Group:    corev1.SchemeGroupVersion.Group,
		Version:  corev1.SchemeGroupVersion.Version,
		Resource: "pods",
	}).Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			t.Log("add pod event")
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			t.Log("update pod event")
		},
		DeleteFunc: func(obj interface{}) {
			t.Log("delete pod event")
		},
	})

	stopper := make(chan struct{})
	defer close(stopper)

	informerFactory.Start(stopper)
	informerFactory.WaitForCacheSync(stopper)
	t.Log("synced")

	<-stopper
}
