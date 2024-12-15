package algorithm

import (
	"fmt"
	"testing"
)

// 1.暴力，循环删除，循环判断 n3
// 2.双指针，n
func incremovableSubarrayCount(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0
	left := 1
	n := len(nums)
	// 从左找符合条件的。意思是右边一个不留
	for left < n && nums[left] > nums[left-1] {
		left++
	}

	res += left
	// 全部移除也算
	if left < n {
		res++
	}

	// 右边留一个到多个
	right := n - 2
	for right >= 0 {
		for left > 0 && nums[left-1] >= nums[right+1] {
			left--
		}

		res += left
		// 没太理解这里，需要再看看
		if left <= right {
			res++
		}
		// 循环停止，因为右边的不满足条件了
		if nums[right] >= nums[right+1] {
			break
		}

		right--
	}

	return res
}

func Test2970(t *testing.T) {
	// 6,5,7,8
	// 1,2,3,4
	nums := []int{1, 2, 3, 4}
	res := incremovableSubarrayCount(nums)
	fmt.Println(res)
}
