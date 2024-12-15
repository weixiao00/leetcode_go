package algorithm

import (
	"fmt"
	"testing"
)

// 动态规划
func canPartition(nums []int) bool {

	if len(nums) == 0 {
		return true
	}

	count := 0
	for _, num := range nums {
		count += num
	}
	// 奇数返回false
	if count&1 == 1 {
		return false
	}
	target := count / 2

	// 二维dp。从i个数里那0个或者多个数之和等于j
	dp := make([][]bool, 0)
	for i := 0; i < len(nums); i++ {
		dp = append(dp, make([]bool, target+1))
		dp[i][0] = true
	}
	// 没什么用，还容易数组越界报错
	//dp[0][nums[0]] = true
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		for j := 1; j <= target; j++ {
			// 大于总和，不选择这个数
			if num > j {
				dp[i][j] = dp[i-1][j]
			} else {
				// 不选择这个数和选择这个数哪个为true。就是true
				dp[i][j] = dp[i-1][j] || dp[i-1][j-num]
			}
		}
	}

	return dp[len(nums)-1][target]
}

func Test416(t *testing.T) {
	nums := []int{9, 1}
	res := canPartition(nums)
	fmt.Println(res)
}
