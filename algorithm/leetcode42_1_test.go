package algorithm

import (
	"fmt"
	"testing"
)

//func trap(height []int) int {
//	if height == nil {
//		return 0
//	}
//	res := 0
//	begin := 0
//	end := len(height) - 1
//	leftMax := height[begin]
//	rightMax := height[end]
//	for begin < end {
//		if leftMax < rightMax {
//			if leftMax > height[begin+1] {
//				res += leftMax - height[begin+1]
//			} else {
//				leftMax = height[begin+1]
//			}
//			begin ++
//		} else {
//			if rightMax > height[end-1] {
//				res += rightMax - height[end-1]
//			} else {
//				rightMax = height[end-1]
//			}
//			end --
//		}
//	}
//	return res
//}

func trap1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0

	leftMax := 0
	rightMax := 0

	left := 0
	right := len(nums) - 1
	for left <= right {
		leftMax = getMax16(leftMax, nums[left])
		rightMax = getMax16(rightMax, nums[right])
		if leftMax <= rightMax {
			//if leftMax > nums[left] {
			res += leftMax - nums[left]
			//}
			left++
		} else if rightMax <= leftMax {
			//if rightMax > nums[right] {
			res += rightMax - nums[right]
			//}
			right--
		}
	}

	return res
}

// 单调栈得算距离
//func trap2(nums []int) int {
//
//	if len(nums) == 0 {
//		return 0
//	}
//
//}

func getMax16(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test1234(t *testing.T) {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap1(nums)
	fmt.Println(res)
}
