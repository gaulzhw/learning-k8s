package linkedlist

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		list     *list
		expected *list
	}{
		{
			list:     newListWithNodes(1, 2, 3, 4, 5),
			expected: newListWithNodes(5, 4, 3, 2, 1),
		},
	}
	for _, test := range tests {
		got := reverse(test.list)
		assert.True(t, reflect.DeepEqual(test.expected, got))
	}
}
