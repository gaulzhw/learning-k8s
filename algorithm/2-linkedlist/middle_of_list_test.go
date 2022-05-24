package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleOfList(t *testing.T) {
	tests := []struct {
		list     *list
		expected int
	}{
		{
			list:     newListWithNodes(1, 2, 3, 4, 5),
			expected: 3,
		},
		{
			list:     newListWithNodes(1, 2, 4, 5),
			expected: 2,
		},
	}
	for _, test := range tests {
		got := test.list.middleOfList()
		assert.Equal(t, test.expected, got)
	}
}
