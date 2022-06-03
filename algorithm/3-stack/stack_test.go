package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := newStack()
	s.push(1)
	s.push(2)
	s.push(3)
	s.pop()
	s.print()
}
