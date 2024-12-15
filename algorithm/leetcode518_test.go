package algorithm

import (
	"fmt"
	"testing"
)

// 二层二维动态规划
func change(amount int, coins []int) int {

	if len(coins) == 0 {
		return 0
	}

	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	for i := 0; i <= len(coins); i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if coins[i-1] <= j {
				dp[i][j] += dp[i][j-coins[i-1]]
			}
		}
	}

	return dp[len(coins)][amount]
}

func Test518(t *testing.T) {
	amount := 5
	coins := []int{1, 2, 5}
	res := change(amount, coins)
	fmt.Println(res)
}
