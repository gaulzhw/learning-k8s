package clientgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInformer(t *testing.T) {
	err := StartInformer()
	assert.NoError(t, err)
}
