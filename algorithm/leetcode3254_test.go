package algorithm

import (
	"fmt"
	"testing"
)

// 解法一：暴力枚举
// 解法二：查个数
// 和3255题是一样的
func resultsArray(nums []int, k int) []int {

	n := len(nums)
	if len(nums) == 0 {
		return []int{}
	}
	res := make([]int, n-k+1)
	for i := range res {
		res[i] = -1
	}

	cnt := 0
	for i := 0; i < n; i++ {
		if i == 0 || nums[i]-nums[i-1] != 1 {
			cnt = 0
		}
		cnt++
		if cnt >= k {
			res[i-k+1] = nums[i]
		}
	}

	return res
}

func Test3254(t *testing.T) {
	nums := []int{1, 2, 3, 4, 3, 2, 5}
	k := 3
	res := resultsArray(nums, k)
	fmt.Println(res)
}
