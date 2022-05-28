package linkedlist

/*
删除链表的倒数第n个节点
给定一个链表，删除链表的倒数第n个节点，并且返回链表的头节点。
使用一趟扫描实现

示例1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
*/

func removeNthFromEnd(l *list, n int) *list {
	// 1. 设置虚拟节点 dummyHead 指向 head
	// 2. 设定双指针 p 和 q，初始都指向虚拟节点 dummyHead
	// 3. 移动 q，直到 p 与 q 之间相隔的元素个数为 n
	// 4. 同时移动 p 与 q，直到 q 指向的为 NULL
	// 5. 将 p 的下一个节点指向下下个节点

	p, q := l.head, l.head
	count := 0
	for q != nil {
		if count <= n {
			q = q.next
			count++
			continue
		}
		p = p.next
		q = q.next
	}
	p.next = p.next.next
	l.length--
	return l
}
