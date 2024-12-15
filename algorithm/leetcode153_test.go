package algorithm

import (
	"fmt"
	"testing"
)

// 二分查找
func findMin(nums []int) int {

	if len(nums) == 0 {
		return -1
	}
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := ((right - left) >> 1) + left
		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		// 必须得要，因为如果没有的话。就不能是mid+1。就是mid
		if mid > 0 && nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func Test153(t *testing.T) {
	nums := []int{4, 5, 1, 2, 3}
	res := findMin(nums)
	fmt.Println(res)
}
