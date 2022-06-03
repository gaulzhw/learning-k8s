package linkedlist

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		input    *list
		k        int
		expected *list
	}{
		{
			input:    newListWithNodes(1, 2, 3, 4, 5),
			k:        2,
			expected: newListWithNodes(2, 1, 4, 3, 5),
		},
	}
	for _, test := range tests {
		got := reverseKGroup(test.input, test.k)
		assert.True(t, reflect.DeepEqual(test.expected, got))
	}
}
