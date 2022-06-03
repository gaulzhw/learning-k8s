package stack

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := newQueue()
	q.push(1)
	q.push(2)
	q.push(3)
	q.pop()
	q.pop()
	q.pop()
	q.print()
}
