package main

import (
	"awesomeProject/tree"
	"encoding/json"
	"fmt"
)

func main() {
	n := &tree.ListNode{
		MinValue: 0,
		MaxValue: 5,
		Head:     nil,
		Next:     nil,
		CurrentNode: &tree.Node{
			Value: 2,
			Next:  nil,
		},
	}

	n.InsertNode(4)
	n.InsertNode(3)
	n.InsertNode(2)
	n.InsertNode(7)
	n.InsertNode(8)
	n.InsertNode(10)
	n.InsertNode(154)
	fmtGetJson(n)
}

func fmtGetJson(a interface{}) {
	marshal, _ := json.Marshal(a)

	fmt.Println(string(marshal))
}
