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
func amountOfTime(root *TreeNode, start int) int {
	if root == nil {
		return 0
	}

	graphMap := make(map[int][]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		val := node.Val
		children := []*TreeNode{node.Left, node.Right}
		for _, child := range children {
			if child == nil {
				continue
			}
			v := child.Val
			// 父加子
			graphMap[val] = append(graphMap[val], v)
			// 子加父
			graphMap[v] = append(graphMap[v], val)
			dfs(child)
		}
	}

	dfs(root)

	time := 0
	queue := list.New()
	queue.PushBack([]int{start, 0})
	// 剪枝
	visited := make(map[int]bool, 0)
	visited[start] = true
	for queue.Len() != 0 {
		// 断言
		arr := queue.Front().Value.([]int)
		// 队列移除头节点
		queue.Remove(queue.Front())
		time = arr[1]
		for _, val := range graphMap[arr[0]] {
			if !visited[val] {
				queue.PushBack([]int{val, time + 1})
			}
			visited[val] = true
		}
	}
	return time
}

func Test2385(t *testing.T) {
	root := TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, nil}}
	time := amountOfTime(&root, 1)
	fmt.Println(time)
}
