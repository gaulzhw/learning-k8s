package linkedlist

/*
相交链表
给定两个单链表的头节点 headA 和 headB，找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回null。

示例1：
输入：intersectVal = 8, listA = [4,1,8,4,5], skipA = 2, skipB = 3
输出：Intersected at '8'
解释：相交节点的值为8（两个链表相交则不能为0）
从各自的表头开始算起，链表A为[4,1,8,4,5]，链表B为[5,6,1,8,4,5]，在A中，相交节点前有2个节点；在B中，相交节点前有3个节点。
*/

// 时间复杂度O(m+n)，空间复杂度O(1)
func intersectionNode(l1 *list, l2 *list) *node {
	// 先让指向长链表的指针多走 na-nb 或 nb-na 步骤
	pA := l1.head.next
	pB := l2.head.next
	if l1.length >= l2.length {
		for i := 0; i < l1.length-l2.length; i++ {
			pA = pA.next
		}
	} else {
		for i := 0; i < l2.length-l1.length; i++ {
			pB = pB.next
		}
	}
	// 同步前进
	for pA != nil && pB != nil && pA != pB {
		pA = pA.next
		pB = pB.next
	}
	if pA == nil || pB == nil {
		return nil
	}
	return pA
}
