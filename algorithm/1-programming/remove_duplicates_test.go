package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input          []int
		expected       []int
		expectedLength int
	}{
		{
			input:          []int{1, 1, 2},
			expected:       []int{1, 2, 2},
			expectedLength: 2,
		},
	}
	for _, test := range tests {
		got := removeDuplicates(test.input)
		assert.Equal(t, test.expected, test.input)
		assert.Equal(t, test.expectedLength, got)
	}
}
