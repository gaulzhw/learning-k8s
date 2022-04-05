package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscoveryClient(t *testing.T) {
	err := InitDiscoveryClient()
	assert.NoError(t, err)
}
