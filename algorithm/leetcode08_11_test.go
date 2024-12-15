package algorithm

import (
	"fmt"
	"testing"
)

// 动态规划
func waysToChange(n int) int {

	if n == 0 {
		return 0
	}

	MOD := 1000000007

	coins := []int{1, 5, 10, 25}
	dp := make([]int, n+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= n; i++ {
			if i >= coin {
				dp[i] = (dp[i] + dp[i-coin]) % MOD
			}
		}
	}

	return dp[n]
}

// 这样会超时，类似三个小朋友分糖果的题。固定一个。算第二个的最大值和最小值，相减+1。因为第一个和第二个固定了。第三个也就固定了
func waysToChange1(n int) int {
	if n == 0 {
		return 0
	}

	MOD := 1000000007
	res := 0
	for i := 0; i*25 <= n; i++ {
		rest := n - 25*i
		for j := 0; j*10 <= rest; j++ {
			res += ((n-i*25-j*10)/5 + 1) % MOD
		}
	}

	return res
}

func Test08_11(t *testing.T) {
	n := 5
	res := waysToChange1(n)
	fmt.Println(res)
}
