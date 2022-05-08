package client

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	kubeconfig = ""
)

func TestClientSet(t *testing.T) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	assert.NoError(t, err)

	client, err := kubernetes.NewForConfig(config)
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestClientSetWithContext(t *testing.T) {
	configLoadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: filepath.Join(homedir.HomeDir(), ".kube", "config")}
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: "kind-hub"}
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(configLoadingRules, configOverrides).ClientConfig()
	assert.NoError(t, err)

	clientset, err := kubernetes.NewForConfig(config)
	assert.NoError(t, err)

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	assert.NoError(t, err)
	for _, node := range nodes.Items {
		t.Logf("node info: %s\n", node.Name)
	}
}

func TestListPods(t *testing.T) {
	clientConfig, err := clientcmd.NewClientConfigFromBytes([]byte(kubeconfig))
	assert.NoError(t, err)
	config, err := clientConfig.ClientConfig()
	assert.NoError(t, err)

	client, err := kubernetes.NewForConfig(config)
	assert.NoError(t, err)

	tests := []struct {
		Name            string
		ResourceVersion string
	}{
		{
			Name:            "List pods with ResourceVersion: 0",
			ResourceVersion: "0",
		},
		{
			Name:            "List pods with empty ResourceVersion",
			ResourceVersion: "",
		},
	}
	for _, test := range tests {
		t.Logf("test for %s\n", test.Name)
		pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{
			ResourceVersion: test.ResourceVersion,
		})
		assert.NoError(t, err)
		for _, pod := range pods.Items {
			t.Logf("%v\t %v\t %v\n", pod.Namespace, pod.Status.Phase, pod.Name)
		}
	}
}
