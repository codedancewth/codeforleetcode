package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}

	b := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  2,
			Next: a,
		},
	}

	c := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  1,
				Next: a,
			},
		},
	}

	// 1 9 1 -> [a]1,4
	//           |
	//      3 2 ->
	fmt.Println(getJson(getIntersectionNode(c, b)))
}

func getJson(a interface{}) string {
	marshal, _ := json.Marshal(a)

	return string(marshal)
}
