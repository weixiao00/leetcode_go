package algorithm

import (
	"fmt"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {

	res := 0
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if k <= 0 {
			return
		}
		preOrder(root.Left)
		k--
		if k == 0 {
			res = root.Val
		}
		preOrder(root.Right)
	}

	preOrder(root)

	return res
}

func Test230(t *testing.T) {
	//root = [3,1,4,null,2], k = 1
	root := &TreeNode{3, nil, nil}
	root1 := &TreeNode{1, nil, nil}
	root2 := &TreeNode{4, nil, nil}
	root3 := &TreeNode{2, nil, nil}

	root.Left = root1
	root.Right = root2

	root1.Right = root3
	res := kthSmallest(root, 2)
	fmt.Println(res)
}
