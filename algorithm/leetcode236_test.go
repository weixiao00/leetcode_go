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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	p1 := false
	q1 := false
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root == p {
			p1 = true
		}
		if root == q {
			q1 = true
		}
		dfs(root.Left)
		dfs(root.Right)
	}

	var res *TreeNode
	var lowestCommon func(root *TreeNode)
	lowestCommon = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root)
		if p1 && q1 {
			res = root
		}
		q1 = false
		p1 = false
		lowestCommon(root.Left)
		lowestCommon(root.Right)
	}

	lowestCommon(root)

	return res
}

func Test236(t *testing.T) {
	//root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
	root3 := &TreeNode{3, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root1 := &TreeNode{1, nil, nil}
	root3.Left = root5
	root3.Right = root1
	root6 := &TreeNode{6, nil, nil}
	root2 := &TreeNode{2, nil, nil}
	root5.Left = root6
	root5.Right = root2
	root0 := &TreeNode{0, nil, nil}
	root8 := &TreeNode{8, nil, nil}
	root1.Left = root0
	root1.Right = root8
	root7 := &TreeNode{7, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root2.Left = root7
	root2.Right = root4
	res := lowestCommonAncestor(root3, root5, root1)
	fmt.Println(res)
}
