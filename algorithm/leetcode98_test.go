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
// 获取某个节点的左右子树的最大和最小值。
// 跟官方题解思路一致。实现稍有区别
func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var getNodeMax func(root *TreeNode) int
	getNodeMax = func(root *TreeNode) int {
		if root == nil {
			return math.MinInt
		}
		leftMax := getNodeMax(root.Left)
		rightMax := getNodeMax(root.Right)

		return getMax(getMax(leftMax, rightMax), root.Val)
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var getNodeMin func(root *TreeNode) int
	getNodeMin = func(root *TreeNode) int {
		if root == nil {
			return math.MaxInt
		}
		leftMax := getNodeMin(root.Left)
		rightMax := getNodeMin(root.Right)

		return getMin(getMin(leftMax, rightMax), root.Val)
	}

	leftMax := getNodeMax(root.Left)
	rightMax := getNodeMin(root.Right)
	return leftMax < root.Val && root.Val < rightMax && isValidBST(root.Left) && isValidBST(root.Right)
}

func Test98(t *testing.T) {
	//root = [5,1,4,null,null,3,6]
	//root := &TreeNode{5, nil, nil}
	//root1 := &TreeNode{1, nil, nil}
	//root2 := &TreeNode{4, nil, nil}
	//root3 := &TreeNode{3, nil, nil}
	//root4 := &TreeNode{6, nil, nil}
	//root.Left = root1
	//root.Right = root2
	//root2.Left = root3
	//root2.Right = root4

	//root := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}

	// [32,26,47,19,null,null,56,null,27]
	root := &TreeNode{32, nil, nil}
	root1 := &TreeNode{26, nil, nil}
	root2 := &TreeNode{47, nil, nil}
	root3 := &TreeNode{19, nil, nil}
	root4 := &TreeNode{56, nil, nil}
	root5 := &TreeNode{27, nil, nil}
	root.Left = root1
	root.Right = root2
	root1.Left = root3
	root3.Right = root5
	root2.Right = root4
	res := isValidBST(root)
	fmt.Println(res)
}
