package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func TestRestClient(t *testing.T) {
	restClient, err := newRestClient()
	assert.NoError(t, err)
	assert.NotNil(t, restClient)

	result := &corev1.PodList{}

	// /api/v1/namespaces/{namespace}/pods
	err = restClient.Get().
		Namespace("kube-system").
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 100}, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	assert.NoError(t, err)
	for _, d := range result.Items {
		t.Logf("%v\t %v\t %v\n", d.Namespace, d.Status.Phase, d.Name)
	}
}
