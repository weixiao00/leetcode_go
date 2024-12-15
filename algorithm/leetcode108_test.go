package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	var subArr func(begin, end int) []int
	subArr = func(begin, end int) []int {
		return nums[begin:end]
	}

	mid := len(nums) / 2
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = sortedArrayToBST(subArr(0, mid))
	root.Right = sortedArrayToBST(subArr(mid+1, len(nums)))
	return root
}
