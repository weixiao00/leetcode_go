package algorithm

import (
	"container/list"
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
//层次遍历
func rightSideView(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	res := make([]int, 0)

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		n := queue.Len()
		for i := 0; i < n; i++ {
			front := queue.Front()
			node := front.Value.(*TreeNode)
			queue.Remove(front)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if i == n-1 {
				res = append(res, node.Val)
			}
		}
	}

	return res
}

func Test199(t *testing.T) {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	res := rightSideView(root)
	fmt.Println(res)
}
