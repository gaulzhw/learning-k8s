package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeNum(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{
			input:    121,
			expected: true,
		},
		{
			input:    -121,
			expected: false,
		},
		{
			input:    10,
			expected: false,
		},
		{
			input:    -101,
			expected: false,
		},
	}
	for _, test := range tests {
		got := isPalindromeNum(test.input)
		assert.Equal(t, test.expected, got)
	}
}
