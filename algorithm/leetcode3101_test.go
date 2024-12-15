package algorithm

import (
	"fmt"
	"testing"
)

func countAlternatingSubarrays(nums []int) int64 {
	if len(nums) == 0 {
		return 0
	}

	res := 1

	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			dp[i] = 1
		} else {
			dp[i] = dp[i-1] + 1
		}
		res += dp[i]
	}

	return int64(res)
}

func Test3101(t *testing.T) {
	nums := []int{0, 1, 1, 1}
	res := countAlternatingSubarrays(nums)
	fmt.Println(res)
}
