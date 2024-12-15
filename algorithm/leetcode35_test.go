package algorithm

import (
	"fmt"
	"testing"
)

// 二分查找
// 返回第一个大于等于target的下标
func searchInsert(nums []int, target int) int {

	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

// 官方题解
// 返回第一个大于等于target的下标
func searchInsert1(nums []int, target int) int {

	left := 0
	rigth := len(nums) - 1

	ans := 0
	for left <= rigth {
		mid := ((rigth - left) >> 1) + left
		if target <= nums[mid] {
			ans = mid
			rigth = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func Test35(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2
	res := searchInsert1(nums, target)
	fmt.Println(res)
}
