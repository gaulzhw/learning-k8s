package client

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/fake"
)

var (
	obj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "group/version",
			"kind":       "TheKind",
			"metadata": map[string]interface{}{
				"namespace": "ns-foo",
				"name":      "name-foo",
			},
		},
	}
)

func TestDynamicClient(t *testing.T) {
	scheme := runtime.NewScheme()

	client := fake.NewSimpleDynamicClient(scheme, obj)
	get, err := client.Resource(schema.GroupVersionResource{
		Group:    "group",
		Version:  "version",
		Resource: "thekinds"},
	).Namespace("ns-foo").Get(context.TODO(), "name-foo", metav1.GetOptions{})
	if err != nil {
		t.Fatal(err)
	}

	expected := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "group/version",
			"kind":       "TheKind",
			"metadata": map[string]interface{}{
				"name":      "name-foo",
				"namespace": "ns-foo",
			},
		},
	}
	assert.True(t, reflect.DeepEqual(get, expected))
}
