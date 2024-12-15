package algorithm

import (
	"fmt"
	"testing"
)

// 双指针
func incremovableSubarrayCount1(nums []int) int64 {
	if len(nums) == 0 {
		return 0
	}

	res := 0
	n := len(nums)
	left := 0

	for left < n-1 && nums[left] < nums[left+1] {
		left++
	}
	// 全部符合条件
	if left == n-1 {
		return int64((n * (n + 1)) / 2)
	}

	res += left + 2
	for right := n - 1; right > 0; right-- {
		if right < n-1 && nums[right] >= nums[right+1] {
			break
		}

		for left >= 0 && nums[left] >= nums[right] {
			left--
		}

		res += left + 2
	}

	return int64(res)
}

func Test2972(t *testing.T) {
	nums := []int{6, 5, 7, 8}
	res := incremovableSubarrayCount1(nums)
	fmt.Println(res)
}
