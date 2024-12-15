package algorithm

import (
	"fmt"
	"testing"
)

// 单调栈
// 不是单调栈的题目
// 没通过
func lengthOfLIS(nums []int) int {

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

	res := 1
	stack := make([]int, 0)
	for i := range nums {
		for len(stack) > 0 && stack[len(stack)-1] >= nums[i] {
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, nums[i])
		res = getMax(res, len(stack))
	}

	return res
}

// 动态规划
func lengthOfLIS1(nums []int) int {

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

	res := 1

	dp := make([]int, len(nums))
	for i := range nums {
		max := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				max = getMax(max, dp[j])
			}
		}
		dp[i] = max + 1
		res = getMax(res, dp[i])
	}

	return res
}

func Test300(t *testing.T) {
	nums := []int{0, 1, 0, 3, 2, 3}
	//nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLIS1(nums)
	fmt.Println(res)
}
