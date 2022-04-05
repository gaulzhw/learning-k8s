package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestClient(t *testing.T) {
	err := InitRestClient()
	assert.NoError(t, err)
}
