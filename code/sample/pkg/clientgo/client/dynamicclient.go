package client

import (
	"context"
	"log"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// https://xinchen.blog.csdn.net/article/details/113795523
func InitDynamicClient() {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return
	}

	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}

	unstructObj, err := dynamicClient.Resource(gvr).
		Namespace("kube-system").
		List(context.TODO(), metav1.ListOptions{Limit: 100})
	if err != nil {
		return
	}

	pods := &corev1.PodList{}

	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructObj.UnstructuredContent(), pods)
	if err != nil {
		return
	}

	log.Printf("namespace\t status\t\t name\n")

	for _, pod := range pods.Items {
		log.Printf("%v\t %v\t %v\n", pod.Namespace, pod.Status.Phase, pod.Name)
	}
}
