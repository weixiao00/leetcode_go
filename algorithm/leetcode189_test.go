package algorithm

import (
	"fmt"
	"testing"
)

// 解法一：移动元素，最简单
// 解法二：反转排序(交换)
func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}

	k %= len(nums)

	var swap func(nums []int, begin, end int)
	swap = func(nums []int, begin, end int) {

		// a = a + b
		// b = a - b
		// a = a - b
		// 异或也可以
		for begin < end {
			nums[begin] = nums[begin] + nums[end]
			nums[end] = nums[begin] - nums[end]
			nums[begin] = nums[begin] - nums[end]
			begin++
			end--
		}
	}

	swap(nums, 0, len(nums)-1)
	swap(nums, 0, k-1)
	swap(nums, k, len(nums)-1)
}

func Test189(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 8)
	fmt.Println(nums)
}
