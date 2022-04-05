package client

import (
	"context"
	"log"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// https://xinchen.blog.csdn.net/article/details/113487087
func InitRestClient() error {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return err
	}

	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		return err
	}

	result := &corev1.PodList{}

	namespace := "kube-system"

	// /api/v1/namespaces/{namespace}/pods
	err = restClient.Get().
		Namespace(namespace).
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 100}, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	if err != nil {
		return err
	}

	log.Printf("namespace\t status\t\t name\n")

	for _, d := range result.Items {
		log.Printf("%v\t %v\t %v\n", d.Namespace, d.Status.Phase, d.Name)
	}
	return nil
}
