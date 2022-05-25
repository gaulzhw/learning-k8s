package linkedlist

/*
存在一个按升序排列的链表，给你链表的头节点 head，请你删除所有重复的元素，使每个元素只出现一次。
返回同样按升序排序的结果链表。
*/

func (l *list) deleteDuplicates() {
	pre := l.head
	p := l.head.next
	for p != nil {
		if p.value != pre.value {
			p, pre = p.next, p
		} else {
			p, pre.next = p.next, p.next
		}
	}
	l.tail = pre
}
