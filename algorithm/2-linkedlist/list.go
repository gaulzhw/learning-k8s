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
