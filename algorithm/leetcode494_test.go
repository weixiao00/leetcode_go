package algorithm

import (
	"fmt"
	"testing"
)

func findTargetSumWays(nums []int, target int) int {

	if len(nums) == 0 {
		return 0
	}

	var dfs func(num, j int) int
	dfs = func(num, j int) int {
		if j == len(nums) {
			if num == target {
				return 1
			}
			return 0
		}
		res := 0
		res += dfs(num+nums[j], j+1)
		res += dfs(num-nums[j], j+1)
		return res
	}

	res := 0
	res += dfs(nums[0], 1)
	res += dfs(-nums[0], 1)

	return res
}

func Test494(t *testing.T) {
	nums := []int{1}
	target := 1
	res := findTargetSumWays(nums, target)
	fmt.Println(res)
}
