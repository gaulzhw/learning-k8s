package linkedlist

/*
两数相加
给定两个非空链表，表示两个非负的整数。每个数字都是按照逆序的方式存储，每个节点只能存储一位数字。
将两个数相加，并以相同形式返回一个表示和的链表。
可以假设除了数字0之外，这两个数都不会以0开头。

示例1：
输入：2->4->3, 5->6->4
输出：7->0->8
*/

// 时间复杂度O(n)，空间复杂度O(n)，n表示两个链表的最大长度
func addTwoNums(l1 *list, l2 *list) *list {
	l := newList()

	p1 := l1.head.next
	p2 := l2.head.next
	carry := 0 // 表示进位
	for p1 != nil || p2 != nil {
		sum := 0
		if p1 != nil {
			sum += p1.value
			p1 = p1.next
		}
		if p2 != nil {
			sum += p2.value
			p2 = p2.next
		}
		if carry != 0 {
			sum += carry
		}
		l.append(sum % 10)
		carry = sum / 10
	}

	if carry != 0 {
		l.append(carry)
	}
	return l
}
