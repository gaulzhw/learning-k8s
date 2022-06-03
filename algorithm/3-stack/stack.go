package stack

import (
	"container/list"
	"fmt"
)

type stack struct {
	*list.List
}

func newStack() *stack {
	return &stack{
		List: list.New(),
	}
}

func (q *stack) push(element any) {
	q.List.PushFront(element)
}

func (q *stack) pop() any {
	val := q.List.Front()
	if val == nil {
		return nil
	}
	q.List.Remove(val)
	return val.Value
}

func (q *stack) print() {
	front := q.List.Front()
	for front != nil {
		fmt.Println(front.Value)
		front = front.Next()
	}
}
