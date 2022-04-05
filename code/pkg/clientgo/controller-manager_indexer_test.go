package clientgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestControllerWithIndex(t *testing.T) {
	err := StartControllerWithIndex()
	assert.NoError(t, err)
}
