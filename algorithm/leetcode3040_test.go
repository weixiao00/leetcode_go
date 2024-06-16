package algorithm

import (
	"fmt"
	"testing"
)

type initOp struct {
	nums []int
	memo [][]int
}

// todo leetcode官方的解法。没太理解。还得再看看
func maxOperations1(nums []int) int {
	res := 0
	if len(nums) == 0 || len(nums) == 1 {
		return res
	}

	var helper func(i, j, target int) int
	helper = func(i, j, target int) int {
		memo := make([][]int, len(nums))
		for i := range memo {
			memo[i] = make([]int, len(nums))
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}

		var dfs func(i, j, target int) int
		dfs = func(i, j, target int) int {
			ans := 0
			if i >= j {
				return ans
			}
			if memo[i][j] != -1 {
				return memo[i][j]
			}

			if nums[i]+nums[j] == target {
				ans = getMax9(ans, dfs(i+1, j-1, target)+1)
			}
			if nums[i]+nums[i+1] == target {
				ans = getMax9(ans, dfs(i+2, j, target)+1)
			}
			if nums[j]+nums[j-1] == target {
				ans = getMax9(ans, dfs(i, j-2, target)+1)
			}
			memo[i][j] = ans

			return ans
		}

		return dfs(i, j, target)
	}

	i := 0
	j := len(nums) - 1
	res = helper(i, j, nums[i]+nums[j])
	res = getMax9(res, helper(i, j, nums[i]+nums[i+1]))
	res = getMax9(res, helper(i, j, nums[j]+nums[j-1]))

	return res
}

func getMax9(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test3040(t *testing.T) {
	nums := []int{3, 2, 1, 2, 3, 4}
	operations1 := maxOperations1(nums)
	fmt.Println(operations1)
}
