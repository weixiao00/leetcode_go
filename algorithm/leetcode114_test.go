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
// 使用的额外的存储空间，可以原地转的
func flatten(root *TreeNode) {

	res := make([]*TreeNode, 0)
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root)
		preOrder(root.Left)
		preOrder(root.Right)
	}

	preOrder(root)

	tmp := root
	for i := 0; i < len(res)-1; i++ {
		tmp.Left = nil
		tmp.Right = res[i+1]
		tmp = tmp.Right
	}
}

// 不使用额外空间，原地转
func flatten1(root *TreeNode) {

	if root == nil {
		return
	}

	var auto func(root *TreeNode)
	auto = func(root *TreeNode) {
		if root == nil {
			return
		}
		auto(root.Left)
		auto(root.Right)
		if root.Left != nil {
			tmp := root.Right
			root.Right = root.Left
			root.Left = nil
			for root.Right != nil {
				root = root.Right
			}
			root.Right = tmp
		}
	}

	auto(root)
}

func Test114(t *testing.T) {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	flatten1(root)
	fmt.Println(root.Val)
}
