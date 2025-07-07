package listNode

// https://leetcode.cn/problems/intersection-of-two-linked-lists/description/?envType=study-plan-v2&envId=top-100-liked
// 相交链表 ,官方题解，看上去很简单，但是要用哈希去重
// get！！！
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := make(map[*ListNode]bool)

	for next := headA; next != nil; next = next.Next { // a的先预存
		a[next] = true
	}

	for next := headB; next != nil; next = next.Next { // b的判断
		if a[next] {
			return next
		}
	}

	return nil
}
