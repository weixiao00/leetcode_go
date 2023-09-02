package algorithm

import (
	"sort"
)

func matrixSum(nums [][]int) int {
	if nums == nil {
		return 0
	}
	m := len(nums)
	n := len(nums[0])
	for i := 0; i < m; i++ {
		sort.Ints(nums[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		maxVal := 0
		for j := 0; j < m; j++ {
			// 没有三元运算
			if nums[j][i] > maxVal {
				maxVal = nums[j][i]
			}
		}

		res += maxVal
	}

	return res
}
