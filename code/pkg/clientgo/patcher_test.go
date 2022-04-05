package clientgo

import (
	"testing"
	// +kubebuilder:scaffold:imports

	"github.com/stretchr/testify/assert"
)

func TestPatch(t *testing.T) {
	err := Patch()
	assert.NoError(t, err)
}
