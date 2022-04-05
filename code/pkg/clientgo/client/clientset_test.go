package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientSet(t *testing.T) {
	err := InitClientSet()
	assert.NoError(t, err)
}
