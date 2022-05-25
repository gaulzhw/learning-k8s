package linkedlist

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1    *list
		list2    *list
		expected *list
	}{
		{
			list1:    newListWithNodes(1, 2, 4),
			list2:    newListWithNodes(1, 3, 4),
			expected: newListWithNodes(1, 1, 2, 3, 4, 4),
		},
	}
	for _, test := range tests {
		got := mergeTwoLists(test.list1, test.list2)
		assert.True(t, reflect.DeepEqual(test.expected, got))
	}
}
