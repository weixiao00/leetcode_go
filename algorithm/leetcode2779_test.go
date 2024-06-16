package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

// 排序+滑动窗口
func maximumBeauty(nums []int, k int) int {

	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	res := 0
	left := 0
	for right := range nums {
		for nums[right]-nums[left] > 2*k {
			left++
		}
		res = getMax14(res, right-left+1)
	}

	return res
}

func getMax14(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test2779(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	k := 10
	beauty := maximumBeauty(nums, k)
	fmt.Println(beauty)
}
