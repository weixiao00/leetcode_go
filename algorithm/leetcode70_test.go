package algorithm

import (
	"fmt"
	"testing"
)

// 动态规划
func climbStairs(n int) int {

	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func Test70(t *testing.T) {
	n := 3
	res := climbStairs(n)
	fmt.Println(res)
}
