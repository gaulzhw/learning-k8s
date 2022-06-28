package cgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCSum(t *testing.T) {
	assert.Equal(t, 2, sum(1, 1))
}
