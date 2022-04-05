package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicClient(t *testing.T) {
	err := InitDynamicClient()
	assert.NoError(t, err)
}
