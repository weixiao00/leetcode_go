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

// 从上到下递归
func pathSum(root *TreeNode, targetSum int) int {

	count := 0
	var dfs func(root *TreeNode, curVal int)
	dfs = func(root *TreeNode, curVal int) {
		if root == nil {
			return
		}
		curVal += root.Val
		if curVal == targetSum {
			count++
		}
		dfs(root.Left, curVal)
		dfs(root.Right, curVal)
		return
	}

	var pathSum func(root *TreeNode)
	pathSum = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root, 0)
		pathSum(root.Left)
		pathSum(root.Right)
	}

	pathSum(root)
	return count
}

func Test437(t *testing.T) {
	//root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
	root := &TreeNode{5, nil, nil}
	root1 := &TreeNode{4, nil, nil}
	root2 := &TreeNode{8, nil, nil}
	root.Left = root1
	root.Right = root2
	root3 := &TreeNode{11, nil, nil}
	root1.Left = root3
	root4 := &TreeNode{13, nil, nil}
	root5 := &TreeNode{4, nil, nil}
	root2.Left = root4
	root2.Right = root5
	root6 := &TreeNode{7, nil, nil}
	root7 := &TreeNode{2, nil, nil}
	root3.Left = root6
	root3.Right = root7
	root8 := &TreeNode{5, nil, nil}
	root9 := &TreeNode{1, nil, nil}
	root5.Left = root8
	root5.Right = root9
	targetSum := 22
	res := pathSum(root, targetSum)
	fmt.Println(res)
}
