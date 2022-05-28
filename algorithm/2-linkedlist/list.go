package linkedlist

type list struct {
	head *node
	tail *node
}

type node struct {
	value int
	next  *node
}

func newList() *list {
	head := &node{
		next: nil,
	}
	return &list{
		head: head,
		tail: head,
	}
}

func newListWithNodes(vals ...int) *list {
	head := &node{
		next: nil,
	}
	p := head
	for _, val := range vals {
		n := &node{
			value: val,
		}
		p.next = n
		p = n
	}
	return &list{
		head: head,
		tail: p,
	}
}

func (l *list) append(val int) {
	n := &node{
		value: val,
	}
	l.tail.next = n
	l.tail = n
}

func (l *list) toSlice() []int {
	result := make([]int, 0)
	p := l.head.next
	for p != nil {
		result = append(result, p.value)
		p = p.next
	}
	return result
}
