package programming

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfTwo(t *testing.T) {
	tests := []struct {
		inputs   []int
		target   int
		expected []int
	}{
		{
			inputs:   []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
	}
	for _, test := range tests {
		got := sumOfTwo(test.inputs, test.target)
		assert.True(t, reflect.DeepEqual(test.expected, got))
	}
}
