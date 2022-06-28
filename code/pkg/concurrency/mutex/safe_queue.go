package mutex

import (
	"sync"
)

type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewSliceQueue(n int) *SliceQueue {
	return &SliceQueue{
		data: make([]interface{}, n),
	}
}

// Enqueue 把值放在队尾
func (q *SliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, v)
}

// Dequeue 移去队头并返回
func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.data) == 0 {
		return nil
	}

	v := q.data[0]
	q.data = q.data[1:]
	return v
}
