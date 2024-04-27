package algorithm

func rangeSumBST(root *TreeNode, low int, high int) int {

	if root == nil {
		return 0
	}

	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}

	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
