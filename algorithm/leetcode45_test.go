package algorithm

import (
	"fmt"
	"testing"
)

// 还是贪心，dfs会超时
// 就是往远跳
func jump(nums []int) int {

	if len(nums) < 2 {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	mostIdx := 0
	end := 0
	for i := 0; i < len(nums); i++ {
		if i <= mostIdx {
			mostIdx = getMax(mostIdx, i+nums[i])
		}
		// 可以跳到最远，步数加1返回
		if mostIdx >= len(nums)-1 {
			return res + 1
		}
		// 上一跳到了最远
		if i == end {
			end = mostIdx
			res++
		}
	}

	return res
}

func Test45(t *testing.T) {
	nums := []int{1, 2, 1, 1, 1}
	rs := jump(nums)
	fmt.Println(rs)
}
