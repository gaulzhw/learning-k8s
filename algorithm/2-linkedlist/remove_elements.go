package linkedlist

/*
给你一个链表的头节点 head 和一个整数 val，删除链表中所有满足 Node.val == val 的节点，并返回新的头节点。

示例1：
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

示例2：
输入：head = [], val = 1
输出：[]

示例3：
输入：head = [7,7,7,7], val = 7
输出：[]
*/

func (l *list) removeElements(val int) {
	pre := l.head
	p := l.head.next
	for p != nil {
		if p.value == val {
			p = p.next
			pre.next = p
		} else {
			pre, p = p, p.next
		}
	}
	l.tail = pre
}
