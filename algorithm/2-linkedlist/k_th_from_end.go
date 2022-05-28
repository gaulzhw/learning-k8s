package linkedlist

/*
链表中倒数第k个节点
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第一个节点。
例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。

示例1：
输入：1->2->3->4->5, k = 2
输出：4->5
*/

func getKthFromEnd(l *list, k int) *node {
	p, q := l.head, l.head
	count := 0
	for q != nil {
		if count < k {
			q = q.next
			count++
			continue
		}
		p = p.next
		q = q.next
	}
	return p
}
