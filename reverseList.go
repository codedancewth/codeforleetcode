package main

// 你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
func reverseList(head *ListNode) *ListNode {
	var node *ListNode
	cur := head

	for cur != nil {
		temp := cur.Next // 第一步设定一个值把
		cur.Next = node
		node = cur
		cur = temp
	}

	return node
}
