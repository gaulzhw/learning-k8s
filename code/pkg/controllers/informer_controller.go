package controllers

import (
	corev1 "k8s.io/api/core/v1"
	appsinformers "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	appslisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
)

const controllerAgentName = "sample-controller"

// Controller is the controller implementation for Foo resources
type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface

	deploymentsLister appslisters.DeploymentLister
	deploymentsSynced cache.InformerSynced

	recorder record.EventRecorder
}

// NewController returns a new sample controller
func NewController(kubeclientset kubernetes.Interface, deploymentInformer appsinformers.DeploymentInformer) *Controller {
	klog.V(4).Info("Creating event broadcaster")

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})

	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset:     kubeclientset,
		deploymentsLister: deploymentInformer.Lister(),
		deploymentsSynced: deploymentInformer.Informer().HasSynced,
		recorder:          recorder,
	}

	klog.Info("Setting up event handlers")
	deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {

		},
		UpdateFunc: func(old, new interface{}) {

		},
		DeleteFunc: func(obj interface{}) {

		},
	})

	return controller
}

// Run waits until policy informer to be synced
func (c *Controller) Run(stopCh <-chan struct{}) {
	if !cache.WaitForCacheSync(stopCh, c.deploymentsSynced) {
		klog.Info("failed to sync informer cache")
		return
	}
	<-stopCh
}
