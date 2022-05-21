package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "42",
			expected: 42,
		},
		{
			input:    "   -42",
			expected: -42,
		},
		{
			input:    "4193 with words",
			expected: 4193,
		},
		{
			input:    "words and 987",
			expected: 0,
		},
		{
			input:    "-91283472332",
			expected: -2147483648,
		},
	}
	for _, test := range tests {
		got := atoi(test.input)
		assert.Equal(t, test.expected, got)
	}
}
