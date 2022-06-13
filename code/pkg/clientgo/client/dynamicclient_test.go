package client

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
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
		Resource: "thekinds",
	}).Namespace("ns-foo").Get(context.TODO(), "name-foo", metav1.GetOptions{})
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

func TestDynamicClientForNamespace(t *testing.T) {
	scheme := runtime.NewScheme()

	unstructuredNS, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Namespace",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	})
	assert.NoError(t, err)

	client := fake.NewSimpleDynamicClient(scheme, &unstructured.Unstructured{Object: unstructuredNS})

	gvr := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "namespaces",
	}
	namespace, err := client.Resource(gvr).Get(context.TODO(), "test", metav1.GetOptions{})
	assert.NoError(t, err)
	t.Logf("%+v", namespace)
}
