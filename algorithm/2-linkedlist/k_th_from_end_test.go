package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKthFromEnd(t *testing.T) {
	tests := []struct {
		input    *list
		n        int
		expected int
	}{
		{
			input:    newListWithNodes(1, 2, 3, 4, 5),
			n:        2,
			expected: 4,
		},
	}
	for _, test := range tests {
		got := getKthFromEnd(test.input, test.n)
		assert.Equal(t, test.expected, got.value)
	}
}
