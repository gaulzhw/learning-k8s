package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftReverseWord(t *testing.T) {
	tests := []struct {
		input    string
		bits     int
		expected string
	}{
		{
			input:    "abcdefg",
			bits:     2,
			expected: "cdefgab",
		},
		{
			input:    "lrloseumgh",
			bits:     6,
			expected: "umghlrlose",
		},
	}
	for _, test := range tests {
		got := leftReverseWord(test.input, test.bits)
		assert.Equal(t, test.expected, got)
	}
}
