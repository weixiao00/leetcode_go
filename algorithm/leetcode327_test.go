package algorithm

import (
	"fmt"
	"testing"
)

// leetcode这么搞超时了
func countRangeSum(nums []int, lower int, upper int) int {

	if len(nums) == 0 {
		return 0
	}

	prefixSum := make([]int, len(nums)+1)

	for i := 1; i <= len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	res := 0
	for i := 1; i < len(prefixSum); i++ {
		for j := 0; j < i; j++ {
			if prefixSum[i]-prefixSum[j] >= lower && prefixSum[i]-prefixSum[j] <= upper {
				res++
			}
		}
	}

	return res
}

func Test327(t *testing.T) {
	nums := []int{2147483647, -2147483648, -1, 0}
	lower := -1
	upper := 0
	res := countRangeSum(nums, lower, upper)
	fmt.Println(res)
}
