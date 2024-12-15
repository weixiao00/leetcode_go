package algorithm

import (
	"fmt"
	"math"
	"testing"
)

// 动态规划
// 没通过，内存不够了
func integerReplacement(n int) int {

	dp := make([]int, n+1)

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 2; i <= n; i++ {
		curMin := math.MaxInt32
		if i&1 == 0 {
			curMin = dp[i/2]
		} else {
			curMin = dp[i-1]
			curMin = getMin(curMin, dp[(i+1)/2]+1)
		}
		dp[i] = curMin + 1
	}

	return dp[n]
}

var mp = make(map[int]int)

// dfs+记忆化搜索
func integerReplacement1(n int) int {

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	res, ok := mp[n]
	if ok {
		return res
	}

	if n == 1 {
		return 0
	}

	min := math.MaxInt32
	if n&1 == 0 {
		min = getMin(min, integerReplacement1(n/2)+1)
	} else {
		min = getMin(min, integerReplacement1(n-1)+1)
		min = getMin(min, integerReplacement1(n+1)+1)
	}
	mp[n] = min

	return min
}

func Test397(t *testing.T) {
	n := 4
	res := integerReplacement1(n)
	fmt.Println(res)
}
