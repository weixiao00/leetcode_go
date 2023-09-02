package algorithm

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		maxValue := arr[i-1]
		for j := i - 1; j >= max(0, i-k); j-- {
			dp[i] = max(dp[i], dp[j]+maxValue*(i-j))
			if j > 0 {
				maxValue = max(maxValue, arr[j-1])
			}
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
