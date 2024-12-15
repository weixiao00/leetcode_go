package algorithm

import (
	"fmt"
	"testing"
)

func countSquares(matrix [][]int) int {

	if len(matrix) == 0 {
		return 0
	}

	dp := make([][]int, 0, len(matrix))

	m := len(matrix)
	n := len(matrix[0])
	// 初始化数组
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}

	res := 0
	// 初始化dp
	for i := 0; i < m; i++ {
		dp[i][0] = matrix[i][0]
		res += dp[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = matrix[0][i]
		res += dp[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			dp[i][j] = getMin5(getMin5(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			res += dp[i][j]
		}
	}

	return res
}

func getMin5(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test1277(t *testing.T) {
	matrix := [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}
	res := countSquares(matrix)
	fmt.Println(res)
}
