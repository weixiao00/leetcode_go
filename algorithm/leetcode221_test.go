package algorithm

import (
	"fmt"
	"testing"
)

// 动态规划
// 正方形面积最大，就是边长最大
func maximalSquare(matrix [][]byte) int {

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
		dp[i][0] = int(matrix[i][0] - '0')
		res = getMax19(res, dp[i][0])
	}
	for i := 0; i < n; i++ {
		dp[0][i] = int(matrix[0][i] - '0')
		res = getMax19(res, dp[0][i])
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			dp[i][j] = getMin4(getMin4(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			res = getMax19(res, dp[i][j])
		}
	}

	return res * res
}

// 暴力解法
func maximalSquare1(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxSide := 0
			if matrix[i][j] == '0' {
				continue
			}
			maxSide = 1
			// 1作为左上角的元素
			curIdx := getMin4(m-i, n-j)

			for k := 1; k < curIdx; k++ {
				if matrix[i+k][j+k] == '0' {
					break
				}
				flag := true
				for t := 0; t < k; t++ {
					if matrix[i+k][j+t] == '0' || matrix[i+t][j+k] == '0' {
						flag = false
						break
					}
				}
				if !flag {
					break
				}
				maxSide++
			}
			res = getMax19(res, maxSide)
		}
	}

	return res * res
}

func getMax19(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin4(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test221(t *testing.T) {
	matrix := [][]byte{{'0', '1'}, {'1', '0'}}
	res := maximalSquare1(matrix)
	fmt.Println(res)
}
