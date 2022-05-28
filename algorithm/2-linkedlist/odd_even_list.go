package linkedlist

/*
奇偶链表
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。这里的奇数节点和偶数节点指的是节点编号的奇偶性，不是节点的值的奇偶性。
尝试使用原地算法完成。空间复杂度应为O(1)，时间复杂度O(nodes)，nodes为节点总数。

示例1：
输入：1->2->3->4->5
输出：1->3->5->2->4

示例2：
输入：2->1->3->5->6->4->7
输出：2->3->6->7->1->5->4
*/

func oddEvenList(l *list) *list {
	oddList := newList()
	evenList := newList()

	p := l.head.next
	count := 1
	for p != nil {
		tmp := p.next
		if count%2 == 1 { // 奇数
			p.next = nil
			oddList.tail.next = p
			oddList.tail = p
		} else { // 偶数
			p.next = nil
			evenList.tail.next = p
			evenList.tail = p
		}
		count++
		p = tmp
	}
	oddList.tail.next = evenList.head.next
	oddList.tail = evenList.tail
	return oddList
}
