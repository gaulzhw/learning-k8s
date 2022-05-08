package informer

import (
	"fmt"
	"log"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func TestDynamicInformer(t *testing.T) {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	assert.NoError(t, err)

	client, err := dynamic.NewForConfig(config)
	assert.NoError(t, err)

	informerFactory := dynamicinformer.NewDynamicSharedInformerFactory(client, time.Second*30)
	informerFactory.ForResource(schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}).Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("add pod event")
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("update pod event")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("delete pod event")
		},
	})

	informerFactory.ForResource(schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "configmaps",
	}).Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("add cm event")
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("update cm event")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("delete cm event")
		},
	})

	stopper := make(chan struct{})
	defer close(stopper)

	informerFactory.Start(stopper)
	informerFactory.WaitForCacheSync(stopper)
	log.Println("synced")

	<-stopper
}
