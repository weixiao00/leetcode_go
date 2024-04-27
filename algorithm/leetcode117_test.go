package algorithm

import (
	"container/list"
	"testing"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 层次遍历
func connect(root *Node) *Node {

	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		for e := queue.Front(); e != nil; e = e.Next() {
			//if e.Value.(*Node).Left {
			//
			//}
		}
	}
	return nil
}

func Test117(t *testing.T) {
	node := Node{}
	node.Val = 1
	connect(&node)
}
