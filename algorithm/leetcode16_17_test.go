package algorithm

import (
	"fmt"
	"testing"
)

func maxSubArray1(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := nums[0]
	max := 0
	if nums[0] > 0 {
		max = nums[0]
	}

	for i := 1; i < len(nums); i++ {
		if max < 0 {
			max = 0
		}
		max += nums[i]
		res = getMax(res, max)
	}

	return res
}

func Test16_17(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray1(nums)
	fmt.Println(res)
}
