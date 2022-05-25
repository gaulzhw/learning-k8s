package linkedlist

/*
合并两个排序的链表
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

// 时间复杂度O(m+n)，空间复杂度O(1)
func mergeTwoLists(l1 *list, l2 *list) *list {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	l := newList()
	p := l.head

	p1 := l1.head.next
	p2 := l2.head.next
	for p1 != nil && p2 != nil {
		if p1.value <= p2.value {
			p.next = p1
			p = p1
			p1 = p1.next
		} else {
			p.next = p2
			p = p2
			p2 = p2.next
		}
	}

	if p1 != nil {
		p.next = p1
		l.tail = l1.tail
	}
	if p2 != nil {
		p.next = p2
		l.tail = l2.tail
	}
	return l
}
