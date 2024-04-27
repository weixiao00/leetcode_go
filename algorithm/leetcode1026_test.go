package algorithm

import "math"

func maxAncestorDiff(root *TreeNode) int {
	return dfsTree(root, root.Val, root.Val)
}

func dfsTree(node *TreeNode, min, max int) int {
	if node == nil {
		return 0
	}
	diff := int(math.Max(math.Abs(float64(min-node.Val)), math.Abs(float64(max-node.Val))))
	min = int(math.Min(float64(min), float64(node.Val)))
	max = int(math.Max(float64(max), float64(node.Val)))
	left := dfsTree(node.Left, min, max)
	right := dfsTree(node.Right, min, max)
	diff = int(math.Max(float64(diff), float64(left)))
	diff = int(math.Max(float64(diff), float64(right)))
	return diff
}
