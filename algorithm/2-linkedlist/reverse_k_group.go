package linkedlist

/*
给定一个链表，每k个节点一组进行翻转，返回翻转后的链表。
k是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是k的整数倍，请将最后剩余的节点保持原有顺序。
进阶：
1. 可以设计一个只使用常数额外空间的算法来解决此问题吗
2. 不能只是单纯的改变节点内部的值，而是需要实际进行节点交换

示例1：
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]
*/

func reverseKGroup(l *list, k int) *list {
	result := newList()
	tail := l.tail
	p := l.head.next
	for p != nil {
		// 分组
		count := 0
		q := p
		for q != nil {
			count++
			if count == k {
				break
			}
			q = q.next
		}
		// 最后一组
		if q == nil {
			result.tail.next = p
			result.tail = tail
			return result
		}
		// reverse
		tmp := p.next
		for p != q {
			revTmp := p.next

			p = revTmp
		}
		p = tmp
	}
	return result
}
