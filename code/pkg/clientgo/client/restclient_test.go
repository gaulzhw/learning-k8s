package client

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	restfake "k8s.io/client-go/rest/fake"
)

func TestRestClient(t *testing.T) {
	restClient := &restfake.RESTClient{
		Client: restfake.CreateHTTPClient(func(request *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("{}")),
			}
			return resp, nil
		}),
		NegotiatedSerializer: scheme.Codecs.WithoutConversion(),
	}

	result := &corev1.PodList{}

	// /api/v1/namespaces/{namespace}/pods
	err := restClient.Get().
		Namespace("kube-system").
		Resource("pods").
		Do(context.TODO()).
		Into(result)
	assert.NoError(t, err)
	for _, d := range result.Items {
		t.Logf("%v\t %v\t %v\n", d.Namespace, d.Status.Phase, d.Name)
	}
}
