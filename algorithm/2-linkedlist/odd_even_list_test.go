package linkedlist

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddEvenList(t *testing.T) {
	tests := []struct {
		list     *list
		expected *list
	}{
		{
			list:     newListWithNodes(1, 2, 3, 4, 5),
			expected: newListWithNodes(1, 3, 5, 2, 4),
		},
		{
			list:     newListWithNodes(2, 1, 3, 5, 6, 4, 7),
			expected: newListWithNodes(2, 3, 6, 7, 1, 5, 4),
		},
	}
	for _, test := range tests {
		got := oddEvenList(test.list)
		assert.True(t, reflect.DeepEqual(test.expected, got))
	}
}
