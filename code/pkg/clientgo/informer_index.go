package clientgo

import (
	"log"
	"path/filepath"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func StartInformerWithIndex() {
	home := homedir.HomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return
	}

	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)
	podInformer := informerFactory.Core().V1().Pods().Informer()
	podInformer.GetIndexer().AddIndexers(cache.Indexers{
		"nodeName": indexByPodNodeName,
	})

	stopChan := make(chan struct{})
	defer close(stopChan)

	informerFactory.Start(stopChan)
	informerFactory.WaitForCacheSync(stopChan)

	podList, err := podInformer.GetIndexer().ByIndex("nodeName", "hub-control-plane")
	if err != nil {
		return
	}
	for _, pod := range podList {
		podObj := pod.(*corev1.Pod)
		log.Println(podObj)
	}
}

func indexByPodNodeName(obj interface{}) ([]string, error) {
	pod, ok := obj.(*corev1.Pod)
	if !ok {
		return []string{}, nil
	}

	if len(pod.Spec.NodeName) == 0 || pod.Status.Phase == corev1.PodSucceeded || pod.Status.Phase == corev1.PodFailed {
		return []string{}, nil
	}

	return []string{pod.Spec.NodeName}, nil
}
