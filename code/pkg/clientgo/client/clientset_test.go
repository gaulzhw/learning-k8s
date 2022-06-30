package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreateWithPrefix(t *testing.T) {
	client := newFakeClientSet()
	assert.NotNil(t, client)

	cm, err := client.CoreV1().ConfigMaps("kube-system").Create(context.TODO(), &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "test-",
		},
		Data: map[string]string{
			"test": "test",
		},
	}, metav1.CreateOptions{})
	assert.NoError(t, err)
	t.Logf("%+v", cm)
}

func TestNamespaceLevelResource(t *testing.T) {
	client := newFakeClientSet()
	assert.NotNil(t, client)

	pod, err := client.CoreV1().Pods("test").Create(context.TODO(), &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}, metav1.CreateOptions{})
	assert.NoError(t, err)
	t.Logf("%+v", pod)
}

func TestClusterLevelResource(t *testing.T) {
	client := newFakeClientSet()
	assert.NotNil(t, client)

	pv, err := client.CoreV1().PersistentVolumes().Create(context.TODO(), &corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
	}, metav1.CreateOptions{})
	assert.NoError(t, err)
	t.Logf("%+v", pv)
}

func TestEvents(t *testing.T) {
	client, err := newClientSet()
	assert.NoError(t, err)

	eventEvents, err := client.EventsV1().Events("").List(context.TODO(), metav1.ListOptions{})
	assert.NoError(t, err)
	for _, event := range eventEvents.Items {
		t.Logf("%v", event)
	}

	coreEvents, err := client.CoreV1().Events("").List(context.TODO(), metav1.ListOptions{})
	assert.NoError(t, err)
	for _, event := range coreEvents.Items {
		t.Logf("%v", event)
	}
}
