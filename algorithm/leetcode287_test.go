package algorithm

import (
	"fmt"
	"testing"
)

// hash
func findDuplicate(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	mp := make(map[int]int)
	for i := range nums {
		mp[nums[i]]++
	}

	for key, val := range mp {
		if val > 1 {
			return key
		}
	}

	return -1
}

// 交换查找
func findDuplicate1(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	for i := range nums {
		if nums[i] == i+1 {
			continue
		}
		// 交换
		for nums[i] != i+1 {
			if nums[i] == nums[nums[i]-1] {
				return nums[i]
			}
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		}
	}

	return -1
}

// 二分查找
// 没十分理解
func findDuplicate2(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	res := 0
	for left <= right {
		mid := ((right - left) >> 1) + left
		cnt := 0
		for i := range nums {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			res = mid
		}
	}

	return res
}

func Test287(t *testing.T) {
	nums := []int{8, 7, 1, 10, 17, 15, 18, 11, 16, 9, 19, 12, 5, 14, 3, 4, 2, 13, 18, 18}
	res := findDuplicate2(nums)
	fmt.Println(res)
}
