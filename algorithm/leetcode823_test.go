package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

// 动态规划
func numFactoredBinaryTrees(arr []int) int {
	if arr == nil {
		return 0
	}

	sort.Ints(arr)
	res := 0
	size := len(arr)
	dp := make([]int, size)
	mod := 1000000007
	for i, val := range arr {
		dp[i] = 1
		left := 0
		right := i - 1
		for ; left <= right; left++ {
			for left <= right && arr[left]*arr[right] > val {
				right--
			}
			if left <= right && arr[left]*arr[right] == val {
				if arr[left] == arr[right] {
					dp[i] += dp[left] * dp[right]
				} else {
					dp[i] += dp[left] * dp[right] * 2
				}
			}
		}

		res += dp[i]
	}
	return res % mod
}

func Test823(t *testing.T) {
	arr := []int{2, 4, 5, 10}
	res := numFactoredBinaryTrees(arr)
	fmt.Println(res)
}
