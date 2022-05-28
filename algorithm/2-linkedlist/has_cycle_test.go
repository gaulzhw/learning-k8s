package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		list     *list
		expected bool
	}{
		{
			list: func() *list {
				result := newList()
				result.append(3)
				start := result.append(2)
				result.append(0)
				end := result.append(-1)
				end.next = start
				return result
			}(),
			expected: true,
		},
	}
	for _, test := range tests {
		got := test.list.hasCycle()
		assert.Equal(t, test.expected, got)
	}
}
