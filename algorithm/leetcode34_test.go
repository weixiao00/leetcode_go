package algorithm

import (
	"fmt"
	"testing"
)

// 二分查找
func searchRange(nums []int, target int) []int {

	res := make([]int, 2)
	res[0] = -1
	res[1] = -1

	if len(nums) == 0 {
		return res
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] < target {
				res[0] = mid
				break
			} else {
				right = mid - 1
			}
		} else if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	left = 0
	right = len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				res[1] = mid
				break
			} else {
				left = mid + 1
			}
		} else if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}

func Test34(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 10
	res := searchRange(nums, target)
	fmt.Println(res)
}
