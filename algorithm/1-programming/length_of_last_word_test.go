package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "Hello World",
			expected: 5,
		},
		{
			input:    " ",
			expected: 0,
		},
	}
	for _, test := range tests {
		got := lengthOfLastWord(test.input)
		assert.Equal(t, test.expected, got)
	}
}
