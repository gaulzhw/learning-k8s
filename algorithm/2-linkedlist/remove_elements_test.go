package linkedlist

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		list     *list
		val      int
		expected *list
	}{
		{
			list:     newListWithNodes(1, 2, 6, 3, 4, 5, 6),
			val:      6,
			expected: newListWithNodes(1, 2, 3, 4, 5),
		},
		{
			list:     newListWithNodes(),
			val:      1,
			expected: newListWithNodes(),
		},
		{
			list:     newListWithNodes(7, 7, 7, 7),
			val:      7,
			expected: newListWithNodes(),
		},
	}
	for _, test := range tests {
		test.list.removeElements(test.val)
		assert.True(t, reflect.DeepEqual(test.expected, test.list))
	}
}
