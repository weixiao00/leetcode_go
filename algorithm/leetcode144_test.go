package algorithm

import (
	"fmt"
	"testing"
)

func preorderTraversal(root *TreeNode) []int {

	res := make([]int, 0)

	var f func(node *TreeNode)

	f = func(root *TreeNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val)
		f(root.Left)
		f(root.Right)
	}

	f(root)

	return res
}

func Test144(t *testing.T) {
	root := &TreeNode{1, nil, nil}
	node1 := &TreeNode{2, nil, nil}
	node2 := &TreeNode{3, nil, nil}
	root.Right = node1
	node1.Left = node2
	res := preorderTraversal(root)
	fmt.Println(res)
}
