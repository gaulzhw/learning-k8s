package informer

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

const (
	// maxRetries is the number of times a service will be retried before it is dropped out of the queue.
	// With the current rate-limiter in use (5ms*2^(maxRetries-1)) the following numbers represent the
	// sequence of delays between successive queuings of a service.
	//
	// 5ms, 10ms, 20ms, 40ms, 80ms, 160ms, 320ms, 640ms, 1.3s, 2.6s, 5.1s, 10.2s, 20.4s, 41s, 82s
	maxRetries = 15
)

type FakeController struct {
	t *testing.T

	podLister  corelisters.PodLister
	podsSynced cache.InformerSynced

	queue            workqueue.RateLimitingInterface
	workerLoopPeriod time.Duration
}

func newFakeController(podInformer coreinformers.PodInformer) *FakeController {
	c := &FakeController{
		queue:            workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "fake"),
		workerLoopPeriod: time.Second,
	}

	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %+v: %v", obj, err))
				return
			}
			c.t.Logf("add event, key: %s", key)
			c.queue.Add(key)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err != nil {
				utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %+v: %v", newObj, err))
				return
			}
			c.t.Logf("update event, key: %s", key)
			c.queue.Add(key)
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %+v: %v", obj, err))
				return
			}
			c.t.Logf("delete event, key: %s", key)
			c.queue.Add(key)
		},
	})
	c.podLister = podInformer.Lister()
	c.podsSynced = podInformer.Informer().HasSynced
	return c
}

func (c *FakeController) Run(ctx context.Context, workers int) {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	c.t.Log("Starting fake controller")
	defer c.t.Log("Shutting down fake controller")

	if !cache.WaitForNamedCacheSync("fake", ctx.Done(), c.podsSynced) {
		return
	}

	for i := 0; i < workers; i++ {
		go wait.UntilWithContext(ctx, c.worker, c.workerLoopPeriod)
	}

	<-ctx.Done()
}

func (c *FakeController) worker(ctx context.Context) {
	for c.processNextWorkItem(ctx) {
	}
}

func (c *FakeController) processNextWorkItem(ctx context.Context) bool {
	eKey, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(eKey)

	err := c.sync(ctx, eKey.(string))
	c.handleErr(err, eKey)

	return true
}

func (c *FakeController) handleErr(err error, key interface{}) {
	if err == nil {
		c.queue.Forget(key)
		return
	}

	ns, name, keyErr := cache.SplitMetaNamespaceKey(key.(string))
	if keyErr != nil {
		c.t.Logf("Failed to split meta namespace cache key, key: %s, err: %v", key, err)
	}

	if c.queue.NumRequeues(key) < maxRetries {
		c.t.Logf("Error syncing endpoints, retrying, service: %s, err: %v", klog.KRef(ns, name), err)
		c.queue.AddRateLimited(key)
		return
	}

	c.t.Logf("Dropping service %q out of the queue: %v", key, err)
	c.queue.Forget(key)
	utilruntime.HandleError(err)
}

func (c *FakeController) sync(ctx context.Context, key string) error {
	c.t.Logf("receive event, key: %s", key)
	return nil
}

func TestInformerWithWorkers(t *testing.T) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	assert.NoError(t, err)

	clientset, err := kubernetes.NewForConfig(config)
	assert.NoError(t, err)

	informerFactory := informers.NewSharedInformerFactory(clientset, 0)
	podInformer := informerFactory.Core().V1().Pods()

	go func() {
		c := newFakeController(podInformer)
		c.t = t
		c.Run(context.TODO(), 1)
	}()

	stopper := make(chan struct{})
	defer close(stopper)
	informerFactory.Start(stopper)
	<-stopper
}
