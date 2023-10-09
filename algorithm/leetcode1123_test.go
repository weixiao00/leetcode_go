package algorithm

import (
	"fmt"
	"testing"
)

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	lca, i := dfsLca(root)
	fmt.Println(i)
	return lca
}

func dfsLca(root *TreeNode) (*TreeNode, int) {

	if root == nil {
		return root, 0
	}

	left, leftDepth := dfsLca(root.Left)
	right, rightDepth := dfsLca(root.Right)
	if leftDepth > rightDepth {
		return left, leftDepth + 1
	}
	if rightDepth > leftDepth {
		return right, rightDepth + 1
	}

	return root, leftDepth + 1
}

func Test1123(t *testing.T) {
	root := &TreeNode{2, nil, nil}
	left := &TreeNode{1, nil, nil}
	right := &TreeNode{3, nil, nil}
	root.Left = left
	root.Right = right
	res := lcaDeepestLeaves(root)
	fmt.Println(res)
}
