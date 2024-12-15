package algorithm

import (
	"fmt"
	"testing"
)

// 二维一层动态规划
func minPathSum(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = getMin(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func Test64(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	res := minPathSum(grid)
	fmt.Println(res)
}
