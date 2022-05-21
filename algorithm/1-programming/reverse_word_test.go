package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			input:    " hello world! ",
			expected: "world! hello",
		},
		{
			input:    "a good    example",
			expected: "example good a",
		},
	}
	for _, test := range tests {
		got := reverseWords(test.input)
		assert.Equal(t, test.expected, got)
	}
}
