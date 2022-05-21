package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindromeString(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			input:    "race a car",
			expected: false,
		},
	}
	for _, test := range tests {
		got := isPalindromeString(test.input)
		assert.Equal(t, test.expected, got)
	}
}
