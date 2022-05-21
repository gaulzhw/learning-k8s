package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "We are happy.",
			expected: "We%20are%20happy.",
		},
	}
	for _, test := range tests {
		got := replaceSpace(test.input)
		assert.Equal(t, test.expected, got)
	}
}
