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
// 递归
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{rootVal, nil, nil}

	var split func(inorder []int, val int) (left, right []int)
	split = func(inorder []int, val int) (left, right []int) {
		for i := 0; i < len(inorder); i++ {
			if inorder[i] == val {
				left = inorder[0:i]
				right = inorder[i+1:]
				return left, right
			}
		}

		return left, right
	}

	left, right := split(inorder, rootVal)
	root.Left = buildTree(preorder[1:len(left)+1], left)
	root.Right = buildTree(preorder[len(preorder)-len(right):], right)

	return root
}

func Test105(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(root.Val)
}
