package algorithm

import (
	"fmt"
	"testing"
)

// 二分查找
func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := ((right - left) >> 1) + left
		if nums[mid] == target {
			return mid
		}
		// 说明在左边
		if nums[mid] >= nums[0] {
			// 同一侧
			if target >= nums[0] && target < nums[mid] {
				right = mid - 1
				// 同一侧或者不同侧两种情况
			} else {
				left = mid + 1
			}
			// 说明在右边
		} else {
			if target <= nums[len(nums)-1] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func Test33(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	res := search(nums, target)
	fmt.Println(res)
}
