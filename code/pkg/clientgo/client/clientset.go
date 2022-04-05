package client

import (
	"context"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// https://xinchen.blog.csdn.net/article/details/113788269
func InitClientSet() error {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	pods, err := client.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{Limit: 100})
	if err != nil {
		return err
	}

	log.Printf("namespace\t status\t\t name\n")

	for _, pod := range pods.Items {
		log.Printf("%v\t %v\t %v\n", pod.Namespace, pod.Status.Phase, pod.Name)
	}
	return nil
}
