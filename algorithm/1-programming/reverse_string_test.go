package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []byte
	}{
		{
			input:    []byte("hello"),
			expected: []byte("olleh"),
		},
		{
			input:    []byte("Hannah"),
			expected: []byte("hannaH"),
		},
	}
	for _, test := range tests {
		reverseString(test.input)
		assert.Equal(t, test.input, test.expected)
	}
}
