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
// bfs层次遍历
func widthOfBinaryTree(root *TreeNode) int {

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

	queue := list.New()
	root.Val = 1
	queue.PushBack(root)

	res := 1
	// 层次遍历计算每一层的宽度
	for queue.Len() > 0 {
		n := queue.Len()
		for i := 0; i < n; i++ {
			front := queue.Front()
			queue.Remove(front)
			node := front.Value.(*TreeNode)

			if node.Left != nil {
				node.Left.Val = 2 * node.Val
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				node.Right.Val = 2*node.Val + 1
				queue.PushBack(node.Right)
			}
		}
		if queue.Len() > 0 {
			res = getMax(res, queue.Back().Value.(*TreeNode).Val-queue.Front().Value.(*TreeNode).Val+1)
		}
	}

	return res
}

func Test662(t *testing.T) {
	//root = [1,3,2,5,3,null,9]
	node1 := &TreeNode{1, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node1.Left = node3
	node1.Right = node2
	node5 := &TreeNode{5, nil, nil}
	node33 := &TreeNode{3, nil, nil}
	node3.Left = node5
	node3.Right = node33
	node9 := &TreeNode{9, nil, nil}
	node2.Right = node9
	res := widthOfBinaryTree(node1)
	fmt.Println(res)
}
