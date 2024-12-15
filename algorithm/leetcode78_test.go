package algorithm

import (
	"fmt"
	"testing"
)

// 二进制n就是子集的个数
func subsets(nums []int) [][]int {

	if len(nums) == 0 {
		return [][]int{}
	}

	res := make([][]int, 0)
	n := len(nums)

	for mask := 0; mask < (1 << n); mask++ {
		sub := make([]int, 0)
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				sub = append(sub, nums[i])
			}
		}

		res = append(res, sub)
	}

	return res
}

func Test78(t *testing.T) {
	nums := []int{1, 2}
	res := subsets(nums)
	fmt.Println(res)
}
