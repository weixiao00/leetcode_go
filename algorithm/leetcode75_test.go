package algorithm

import (
	"fmt"
	"testing"
)

// 交换.双指针
func sortColors(nums []int) {

	if len(nums) == 0 {
		return
	}

	left := 0
	right := len(nums) - 1

	// 0和1，2交换
	for left < right {
		for left < len(nums) && nums[left] == 0 {
			left++
		}
		for right >= 0 && nums[right] != 0 {
			right--
		}
		if left >= right {
			break
		}
		tmp := nums[left]
		nums[left] = nums[right]
		nums[right] = tmp
	}

	// 寻找0
	left = 0
	for i := range nums {
		if nums[i] == 0 {
			left++
		}
	}

	right = len(nums) - 1
	for left < right {
		for left < len(nums) && nums[left] == 1 {
			left++
		}
		for right >= 0 && nums[right] == 2 {
			right--
		}
		if left >= right {
			break
		}
		tmp := nums[left]
		nums[left] = nums[right]
		nums[right] = tmp
	}
}

// 单指针
func sortColors1(nums []int) {

	if len(nums) == 0 {
		return
	}

	ptr := 0
	for i := range nums {
		if nums[i] == 0 {
			tmp := nums[ptr]
			nums[ptr] = nums[i]
			nums[i] = tmp
			ptr++
		}
	}

	for i := ptr; i < len(nums); i++ {
		if nums[i] == 1 {
			tmp := nums[ptr]
			nums[ptr] = nums[i]
			nums[i] = tmp
			ptr++
		}
	}
}

func Test75(t *testing.T) {
	nums := []int{0, 0}
	sortColors1(nums)
	for i := range nums {
		fmt.Println(nums[i])
	}
}
