package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/restmapper"
)

func TestDiscoveryClient(t *testing.T) {
	client, err := newFakeDiscoveryClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestMemCachedDiscoveryClient(t *testing.T) {
	client, err := newDiscoveryClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)

	cachedClient := memory.NewMemCacheClient(client)
	groups, err := cachedClient.ServerGroups()
	assert.NoError(t, err)
	for _, group := range groups.Groups {
		t.Log(group.Name)
	}
}

func TestCachedDiscoveryClient(t *testing.T) {
	client, err := newDiscoveryClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)

	// discovery cache client
	cachedDiscoClient := NewMemCacheClient(client)
	restMapper := restmapper.NewDeferredDiscoveryRESTMapper(cachedDiscoClient)

	mapping, err := restMapper.RESTMapping(corev1.SchemeGroupVersion.WithKind("Pod").GroupKind(), corev1.SchemeGroupVersion.Version)
	assert.NoError(t, err)
	t.Logf("%+v", mapping)
}
