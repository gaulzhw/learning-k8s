package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidIP(t *testing.T) {
	tests := []struct {
		inputs   []byte
		expected []byte
	}{
		{
			inputs:   []byte("1.1.1.1"),
			expected: []byte("1[.]1[.]1[.]1"),
		},
		{
			inputs:   []byte("255.100.50.0"),
			expected: []byte("255[.]100[.]50[.]0"),
		},
	}
	for _, test := range tests {
		got := invalidIP(test.inputs)
		assert.Equal(t, test.expected, got)
	}
}
