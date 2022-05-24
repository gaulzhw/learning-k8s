package linkedlist

/*
给定一个头节点为 head 的非空单链表，返回链表中的中间节点。
如果有两个中间节点，则返回第二个中间节点。

示例1：
输入：[1,2,3,4,5]
输出：此列表中的节点3
*/

func (l *list) middleOfList() int {
	slow := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow.value
}
