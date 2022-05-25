package linkedlist

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNums(t *testing.T) {
	tests := []struct {
		list1    *list
		list2    *list
		expected *list
	}{
		{
			list1:    newListWithNodes(2, 4, 3),
			list2:    newListWithNodes(5, 6, 4),
			expected: newListWithNodes(7, 0, 8),
		},
	}
	for _, test := range tests {
		got := addTwoNums(test.list1, test.list2)
		assert.True(t, reflect.DeepEqual(test.expected, got))
	}
}
