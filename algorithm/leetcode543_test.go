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
// 可以不经过根节点
func diameterOfBinaryTree(root *TreeNode) int {

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

	res := 0
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {

		if node == nil {
			return 0
		}

		leftMax := depth(node.Left)
		rightMax := depth(node.Right)
		res = getMax(res, leftMax+rightMax)

		return getMax(leftMax, rightMax) + 1
	}
	depth(root)

	return res
}

func Test543(t *testing.T) {
	//root = [1,2,3,4,5]
	root := &TreeNode{1, nil, nil}
	root2 := &TreeNode{2, nil, nil}
	root3 := &TreeNode{3, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root.Left = root2
	root.Right = root3
	root2.Left = root4
	root2.Right = root5
	res := diameterOfBinaryTree(root)
	fmt.Println(res)
}
