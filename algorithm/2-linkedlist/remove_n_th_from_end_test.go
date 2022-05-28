package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		input    *list
		n        int
		expected *list
	}{
		{
			input:    newListWithNodes(1, 2, 3, 4, 5),
			n:        2,
			expected: newListWithNodes(1, 2, 3, 5),
		},
	}
	for _, test := range tests {
		got := removeNthFromEnd(test.input, test.n)
		assert.Equal(t, test.expected, got)
	}
}
