package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		list     *list
		expected bool
	}{
		{
			list:     newListWithNodes(1, 2),
			expected: false,
		},
		{
			list:     newListWithNodes(1, 2, 2, 1),
			expected: false,
		},
	}
	for _, test := range tests {
		got := test.list.isPalindrome()
		assert.Equal(t, test.expected, got)
	}
}

func TestHalfOfList(t *testing.T) {
	tests := []struct {
		list     *list
		expected int
	}{
		{
			list:     newListWithNodes(1, 2),
			expected: 2,
		},
		{
			list:     newListWithNodes(1, 2, 2, 1),
			expected: 2,
		},
		{
			list:     newListWithNodes(1, 2, 3, 2, 1),
			expected: 3,
		},
		{
			list:     newListWithNodes(1),
			expected: 1,
		},
	}
	for _, test := range tests {
		got := test.list.halfOfList()
		assert.Equal(t, test.expected, got.value)
	}
}
