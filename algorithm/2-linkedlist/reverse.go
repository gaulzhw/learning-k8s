package linkedlist

/*
反转一个单链表

示例：
输入：1->2->3->4->5->NULL
输出：5->4->3->2->1->NULL

进阶：可以迭代或递归地反转链表
*/

func reverse(l *list) *list {
	result := newList()
	p := l.head.next
	result.tail = l.head.next
	for p != nil {
		p, p.next, result.head.next = p.next, result.head.next, p
	}
	return result
}
