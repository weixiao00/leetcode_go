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
func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := depth(node.Left)
		rightDepth := depth(node.Right)

		return getMax(leftDepth, rightDepth) + 1
	}

	return depth(root)
}

func Test104(t *testing.T) {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	res := maxDepth(root)
	fmt.Println(res)
}
