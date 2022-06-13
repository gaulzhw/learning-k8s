package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/discovery/cached/memory"
)

func TestDiscoveryClient(t *testing.T) {
	client, err := newFakeDiscoveryClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestCachedDiscoveryClient(t *testing.T) {
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
