package stack

import (
	"container/list"
	"fmt"
)

type queue struct {
	*list.List
}

func newQueue() *queue {
	return &queue{
		List: list.New(),
	}
}

func (q *queue) push(element any) {
	q.List.PushBack(element)
}

func (q *queue) pop() any {
	val := q.List.Front()
	if val == nil {
		return nil
	}
	q.List.Remove(val)
	return val.Value
}

func (q *queue) print() {
	front := q.List.Front()
	for front != nil {
		fmt.Println(front.Value)
		front = front.Next()
	}
}
