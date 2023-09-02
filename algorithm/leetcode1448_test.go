package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfsG(root, math.MinInt)
}

func dfsG(node *TreeNode, pathMax int) int {
	if node == nil {
		return 0
	}

	res := 0
	if node.Val >= pathMax {
		res++
		pathMax = node.Val
	}
	res += dfsG(node.Left, pathMax) + dfsG(node.Right, pathMax)
	return res
}

func Test1448(t *testing.T) {
	leftL := &TreeNode{3, nil, nil}
	left := &TreeNode{1, leftL, nil}
	rightL := &TreeNode{1, nil, nil}
	rightR := &TreeNode{5, nil, nil}
	right := &TreeNode{4, rightL, rightR}
	root := &TreeNode{3, left, right}
	nodes := goodNodes(root)
	fmt.Println(nodes)
}
