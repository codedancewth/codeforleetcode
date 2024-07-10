package tree

// ListNode 链表的节点来存储槽slot
type ListNode struct {
	Index       int64
	MinValue    int64     // 最低值 0
	MaxValue    int64     // 最高值 5
	Head        *ListNode // 上一条, 头部
	Next        *ListNode // 下一条，尾部
	CurrentNode *Node     // 当前节点的值
}

// Node 树的节点
type Node struct {
	Value int64
	Head  *Node // 上一条, 头部
	Next  *Node // 下一条，尾部
}

// InsertNode 新增结点
func (l *ListNode) InsertNode(value int64) {
	// step1 如果结点为空
	if l == nil {
		l = &ListNode{
			Index:       value / 5,       // 索引
			MinValue:    (value/5)*5 + 1, // 最小的
			MaxValue:    (value/5)*5 + 5, // 最大的
			Head:        nil,
			Next:        nil,
			CurrentNode: nil,
		}
		return
	}

	flag := false // 是否已经完成插入
	for !flag {
		if l.MaxValue < value { // 如果当前的值大于最高槽值，去下一个数据
			if l.Next == nil {
				l.Next = &ListNode{
					Index:    value / 5,
					MinValue: (value/5)*5 + 1,
					MaxValue: (value/5)*5 + 5,
					Head:     nil,
					Next:     nil,
					CurrentNode: &Node{
						Value: value,
						Head:  nil,
						Next:  nil,
					},
				}
				flag = true
				break
			} else {
				l = l.Next
			}
			continue
		}

		if l.MinValue > value {
			if l.Head == nil {
				l.Head = &ListNode{
					Index:    value / 5,
					MinValue: (value/5)*5 + 1,
					MaxValue: (value/5)*5 + 5,
					Head:     nil,
					Next:     nil,
					CurrentNode: &Node{
						Value: value,
						Head:  nil,
						Next:  nil,
					},
				}
				flag = true
				break
			} else {
				l = l.Head
			}
			continue
		}

		for l.CurrentNode != nil {
			if l.CurrentNode.Value == value { // 节点值相等 返回
				flag = true
				break
			}

			if l.CurrentNode.Value < value { // 如果存在值关系，小于的
				if l.CurrentNode.Next == nil {
					l.CurrentNode.Next = &Node{
						Value: value,
						Head:  nil,
						Next:  nil,
					}
					flag = true
					break
				} else {
					l.CurrentNode = l.CurrentNode.Next
				}
				continue
			}

			if l.CurrentNode.Value > value {
				if l.CurrentNode.Head == nil {
					node := &Node{
						Value: value,
						Head:  nil,
						Next:  l.CurrentNode,
					}

					l.CurrentNode = node

					flag = true
					break
				} else {
					l.CurrentNode = l.CurrentNode.Head
				}
				continue
			}
		}
	}
}

func TestBtree() {
	n := &ListNode{
		MinValue: 0,
		MaxValue: 5,
		Head:     nil,
		Next:     nil,
		CurrentNode: &Node{
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
}
