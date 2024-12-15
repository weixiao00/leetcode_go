package algorithm

import (
	"fmt"
	"testing"
)

// 动态规划
func maxProfit3(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	buy1 := -prices[0]
	sell1 := 0
	buy2 := -prices[0]
	sell2 := 0

	for i := 1; i < len(prices); i++ {
		buy1 = getMax(buy1, -prices[i])
		sell1 = getMax(sell1, buy1+prices[i])
		buy2 = getMax(buy2, sell1-prices[i])
		sell2 = getMax(sell2, buy2+prices[i])
	}

	return sell2
}

func Test123(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	res := maxProfit3(prices)
	fmt.Println(res)
}
