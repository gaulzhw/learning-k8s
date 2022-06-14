package informer

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func TestInformer(t *testing.T) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	assert.NoError(t, err)

	clientset, err := kubernetes.NewForConfig(config)
	assert.NoError(t, err)

	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)

	// informer for deployment
	deployInformer := informerFactory.Apps().V1().Deployments()
	informer := deployInformer.Informer()
	deployLister := deployInformer.Lister()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			deploy := obj.(*appsv1.Deployment)
			t.Log("add a deployment:", deploy.Name)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldDeploy := oldObj.(*appsv1.Deployment)
			newDeploy := oldObj.(*appsv1.Deployment)
			t.Log("update deployment:", oldDeploy.Name, newDeploy.Name)
		},
		DeleteFunc: func(obj interface{}) {
			deploy := obj.(*appsv1.Deployment)
			t.Log("delete a deployment:", deploy.Name)
		},
	})

	stopper := make(chan struct{})
	defer close(stopper)

	informerFactory.Start(stopper)
	informerFactory.WaitForCacheSync(stopper)
	t.Log("synced")

	deployments, err := deployLister.Deployments("default").List(labels.Everything())
	assert.NoError(t, err)
	for idx, deploy := range deployments {
		t.Logf("%d -> %s\n", idx+1, deploy.Name)
	}
	<-stopper
}

func TestInformerWithIndex(t *testing.T) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	assert.NoError(t, err)

	clientset, err := kubernetes.NewForConfig(config)
	assert.NoError(t, err)

	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)
	podInformer := informerFactory.Core().V1().Pods().Informer()
	podInformer.GetIndexer().AddIndexers(cache.Indexers{
		"nodeName": func(obj interface{}) ([]string, error) {
			pod, ok := obj.(*corev1.Pod)
			if !ok {
				return []string{}, nil
			}

			if len(pod.Spec.NodeName) == 0 || pod.Status.Phase == corev1.PodSucceeded || pod.Status.Phase == corev1.PodFailed {
				return []string{}, nil
			}

			return []string{pod.Spec.NodeName}, nil
		},
	})

	//eventBroadcaster := record.NewBroadcaster()
	//eventBroadcaster.StartStructuredLogging(0)
	//eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{
	//	Interface: clientset.CoreV1().Events(""),
	//})
	//recorder := eventBroadcaster.NewRecorder(scheme, corev1.EventSource{Component: "pod-informer"})

	stopChan := make(chan struct{})
	defer close(stopChan)

	informerFactory.Start(stopChan)
	informerFactory.WaitForCacheSync(stopChan)

	podList, err := podInformer.GetIndexer().ByIndex("nodeName", "hub-control-plane")
	assert.NoError(t, err)

	for _, pod := range podList {
		podObj := pod.(*corev1.Pod)
		t.Log(podObj)
	}
}

func TestInformerWithLabelSelector(t *testing.T) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	assert.NoError(t, err)

	clientset, err := kubernetes.NewForConfig(config)
	assert.NoError(t, err)

	informerFactory := informers.NewSharedInformerFactoryWithOptions(clientset, 0, informers.WithTweakListOptions(func(opts *v1.ListOptions) {
		opts.LabelSelector = labels.SelectorFromSet(map[string]string{"test": "test"}).String()
	}))

	podInformer := informerFactory.Core().V1().Pods().Informer()
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			t.Log("add event")
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			t.Log("update event")
		},
		DeleteFunc: func(obj interface{}) {
			t.Log("delete event")
		},
	})

	stopChan := make(chan struct{})
	defer close(stopChan)

	informerFactory.Start(stopChan)
	informerFactory.WaitForCacheSync(stopChan)
}
