package algorithm

import (
	"fmt"
	"math"
	"testing"
)

// 自己的解法
func numSquares(num int) int {
	if num == 0 {
		return 0
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([]int, num+1)
	dp[1] = 1
	for i := 2; i <= num; i++ {
		// 循环一半就行，否则会超时
		split := i/2 + 1
		cur := math.MaxInt32
		for j := 1; j < split; j++ {
			if j*j == i {
				cur = 1
				break
			}
			cur = getMin(cur, dp[j]+dp[i-j])
		}
		dp[i] = cur
	}
	return dp[num]
}

// 官方题解
// 一维二层动态规划
func numSquares1(num int) int {

	if num == 0 {
		return 0
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([]int, num+1)
	dp[1] = 1

	for i := 2; i <= num; i++ {
		min := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			min = getMin(min, dp[i-j*j])
		}
		dp[i] = min + 1
	}

	return dp[num]
}

func TestName1(t *testing.T) {
	res := numSquares(12)
	fmt.Println(res)
}
