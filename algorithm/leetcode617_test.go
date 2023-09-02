package algorithm

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	node := &TreeNode{root1.Val + root2.Val, nil, nil}
	node.Left = mergeTrees(root1.Left, root2.Left)
	node.Right = mergeTrees(root1.Right, root2.Right)

	return node
}
