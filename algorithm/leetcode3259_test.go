package algorithm

import (
	"fmt"
	"testing"
)

// 二维动态规划
func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {

	n := len(energyDrinkA)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp[1][0] = energyDrinkA[0]
	dp[1][1] = energyDrinkB[0]

	for i := 2; i <= n; i++ {
		dp[i][0] = getMax(dp[i-1][0], dp[i-2][1]) + energyDrinkA[i-1]
		dp[i][1] = getMax(dp[i-1][1], dp[i-2][0]) + energyDrinkB[i-1]
	}

	return int64(getMax(dp[n][0], dp[n][1]))
}

func Test3259(t *testing.T) {
	energyDrinkA := []int{4, 1, 1}
	energyDrinkB := []int{1, 1, 3}
	res := maxEnergyBoost(energyDrinkA, energyDrinkB)
	fmt.Println(res)
}
