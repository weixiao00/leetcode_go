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
// 遍历树
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil || subRoot == nil {
		return false
	}

	var same func(root, subRoot *TreeNode) bool
	same = func(root, subRoot *TreeNode) bool {
		if root == nil && subRoot == nil {
			return true
		}
		if root == nil && subRoot != nil {
			return false
		}
		if root != nil && subRoot == nil {
			return false
		}
		if root.Val != subRoot.Val {
			return false
		}
		return same(root.Left, subRoot.Left) && same(root.Right, subRoot.Right)
	}

	return same(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func Test572(t *testing.T) {
	//root = [3,4,5,1,2], subRoot = [4,1,2]
	root := &TreeNode{3, nil, nil}
	left1 := &TreeNode{4, nil, nil}
	right1 := &TreeNode{5, nil, nil}
	root.Left = left1
	root.Right = right1
	left2 := &TreeNode{1, nil, nil}
	right2 := &TreeNode{2, nil, nil}
	left1.Left = left2
	left1.Right = right2

	subRoot := &TreeNode{4, nil, nil}
	left := &TreeNode{1, nil, nil}
	right := &TreeNode{2, nil, nil}
	subRoot.Left = left
	subRoot.Right = right

	res := isSubtree(root, subRoot)
	fmt.Println(res)
}
