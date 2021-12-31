package client

import (
	"context"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// https://xinchen.blog.csdn.net/article/details/113788269
func InitClientSet() {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return
	}

	pods, err := clientset.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{Limit: 100})
	if err != nil {
		return
	}

	fmt.Printf("namespace\t status\t\t name\n")

	for _, pod := range pods.Items {
		fmt.Printf("%v\t %v\t %v\n", pod.Namespace, pod.Status.Phase, pod.Name)
	}
}
