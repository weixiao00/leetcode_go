package algorithm

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// 1. 贪心。没通过
// coins := []int{2，5}
//	amount := 6  这种我就通过不了返回了-1
func coinChange(coins []int, amount int) int {

	if len(coins) == 0 || amount == 0 {
		return 0
	}

	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	if amount < coins[len(coins)-1] {
		return -1
	}

	res := 0
	for _, coin := range coins {
		for amount >= coin {
			res++
			amount -= coin
		}
	}

	if amount > 0 {
		return -1
	}

	return res
}

// 动态规划
// 两层一维动态规划
func coinChange1(coins []int, amount int) int {

	if len(coins) == 0 || amount == 0 {
		return 0
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		cur := math.MaxInt32
		for _, coin := range coins {
			if coin <= i {
				cur = getMin(cur, dp[i-coin]+1)
			}
		}
		dp[i] = cur
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

// 递归(记忆化搜索)
func coinChange2(coins []int, amount int) int {

	if len(coins) == 0 {
		return 0
	}

	count := make([]int, amount)
	var coinChange func(rem int) int
	coinChange = func(rem int) int {
		if rem < 0 {
			return -1
		}
		if rem == 0 {
			return 0
		}
		if count[rem-1] != 0 {
			return count[rem-1]
		}

		min := math.MaxInt32
		for i := range coins {
			res := coinChange(rem - coins[i])
			if res >= 0 && res < min {
				min = res + 1
			}
		}
		if min == math.MaxInt32 {
			count[rem-1] = -1
		} else {
			count[rem-1] = min
		}
		return count[rem-1]
	}

	return coinChange(amount)
}

func Test322(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	res := coinChange2(coins, amount)
	fmt.Println(res)
}
