package algorithm

// 贪心.每天都有利润
func maxProfit1(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

// 动态规划
func maxProfit2(prices []int) int {

	dp := make([][]int, len(prices))
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

	// 第一天
	// 没有股票的利润
	dp[0][0] = 0
	// 有股票的利润
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		// 前一天没有股票，前一天有股票今天卖出
		dp[i][0] = getMax(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = getMax(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}
