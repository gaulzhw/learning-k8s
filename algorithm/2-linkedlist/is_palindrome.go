package linkedlist

/*
判断一个链表是否是回文链表。

示例1：
输入：1->2
输出：false

示例2：
输入：1->2->2->1
输出：true
*/

// 时间复杂度O(n)，空间复杂度O(1)
func (l *list) isPalindrome() bool {

	return true
}

// 此时如果是奇数 slow 在链表中间
// 如果是偶数，slow 在中间 n/2+1 那里
func (l *list) halfOfList() *node {
	fast := l.head.next
	slow := l.head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
