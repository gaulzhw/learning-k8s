package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscoveryClient(t *testing.T) {
	client, err := newFakeDiscoveryClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
}
