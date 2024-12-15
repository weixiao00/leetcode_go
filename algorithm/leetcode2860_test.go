package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func countWays(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)
	res := 0
	for i := 0; i <= n; i++ {
		// 选择i个人
		if i > 0 && nums[i-1] >= i {
			continue
		}
		if i < n && nums[i] <= i {
			continue
		}

		res++
	}

	return res
}

func Test2860(t *testing.T) {
	nums := []int{1, 1}
	res := countWays(nums)
	fmt.Println(res)
}
