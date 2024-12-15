package algorithm

import (
	"fmt"
	"math"
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
func maxPathSum(root *TreeNode) int {

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

	res := math.MinInt32
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		cur := node.Val
		leftMax := getMax(0, dfs(node.Left))
		rightMax := getMax(0, dfs(node.Right))

		cur += leftMax + rightMax
		res = getMax(res, cur)

		return node.Val + getMax(leftMax, rightMax)
	}
	dfs(root)

	return res
}

func Test124(t *testing.T) {
	//root = [-10,9,20,null,null,15,7]
	root10 := &TreeNode{-10, nil, nil}
	root9 := &TreeNode{9, nil, nil}
	root20 := &TreeNode{20, nil, nil}
	root15 := &TreeNode{15, nil, nil}
	root7 := &TreeNode{7, nil, nil}
	root10.Left = root9
	root10.Right = root20
	root20.Left = root15
	root20.Right = root7
	res := maxPathSum(root10)
	fmt.Println(res)
}
