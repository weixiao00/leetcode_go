package algorithm

import (
	"fmt"
	"testing"
)

// 动态规划
func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 不偷
	dp[0][0] = 0
	// 偷
	dp[0][1] = nums[0]

	// dp[i][0/1]表示第i家偷或者不偷的最大值
	for i := 1; i < len(nums); i++ {
		dp[i][1] = getMax(dp[i-1][0]+nums[i], dp[i-1][1])
		dp[i][0] = getMax(dp[i-1][0], dp[i-1][1])
	}

	return getMax(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

// 官方题解
func rob1(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([]int, len(nums))

	// dp[i]表示第i家偷或者不偷的最大值
	dp[0] = nums[0]
	dp[1] = getMax(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = getMax(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(nums)-1]
}

func Test198(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	res := rob1(nums)
	fmt.Println(res)
}
