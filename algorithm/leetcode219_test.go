package algorithm

import (
	"fmt"
	"testing"
)

// 重复元素，滑动窗口
func containsNearbyDuplicate(nums []int, k int) bool {

	if len(nums) == 0 {
		return true
	}

	mp := make(map[int]bool, 0)
	for i, num := range nums {
		if mp[num] {
			return true
		}
		mp[num] = true

		if i >= k {
			delete(mp, nums[i-k])
		}
	}

	return false
}

func Test219(t *testing.T) {
	nums := []int{1, 0, 1, 1}
	k := 1
	res := containsNearbyDuplicate(nums, k)
	fmt.Println(res)
}
