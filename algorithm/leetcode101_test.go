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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var symmetric func(left *TreeNode, right *TreeNode) bool
	symmetric = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil && right != nil {
			return false
		}
		if left != nil && right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return symmetric(left.Left, right.Right) && symmetric(right.Left, left.Right)
	}

	return symmetric(root.Left, root.Right)
}

func Test101(t *testing.T) {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{2, &TreeNode{5, nil, nil}, &TreeNode{4, nil, nil}}}
	res := isSymmetric(root)
	fmt.Println(res)
}
