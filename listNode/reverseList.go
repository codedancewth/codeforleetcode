package listNode

// 你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
// 双指针写法
func reverseList(head *ListNode) *ListNode {
	// 向前一位的指针
	var preNode *ListNode
	// 指针指向当前的节点
	cur := head

	for cur != nil {
		temp := cur.Next   // 第一步 先保留原来的指针情况，防止指针链接下一个元素断掉
		cur.Next = preNode // 让cur的指针的下一个元素指向pre node节点
		preNode = cur      // pre的节点指向当前的cur指针
		cur = temp         // 让当前的cur往下继续移动
	}

	return preNode
}
