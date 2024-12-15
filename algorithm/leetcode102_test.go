package algorithm

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	res := make([][]int, 0)

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		n := queue.Len()
		inner := make([]int, 0)
		for i := 0; i < n; i++ {
			front := queue.Front()
			queue.Remove(front)
			node := front.Value.(*TreeNode)
			inner = append(inner, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, inner)
	}

	return res
}
