package clientgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientWithContext(t *testing.T) {
	err := ClientWithContext()
	assert.NoError(t, err)
}
