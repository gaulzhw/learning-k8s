package linkedlist

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		list     *list
		expected *list
	}{
		{
			list:     newListWithNodes(1, 1, 2),
			expected: newListWithNodes(1, 2),
		},
	}
	for _, test := range tests {
		test.list.deleteDuplicates()
		assert.True(t, reflect.DeepEqual(test.expected, test.list))
	}
}
