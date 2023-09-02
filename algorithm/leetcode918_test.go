package algorithm

import (
	"fmt"
	"testing"
)

// 子数组就是连续的
func maxSubarraySumCircular(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	n := len(nums)
	// 两边开始搜索(头尾是连在一起的)
	leftMax := make([]int, 0, n)
	curVale := nums[0]
	resValue := nums[0]
	leftMax = append(leftMax, curVale)
	leftSum := nums[0]
	// 从前向后检索
	for i := 1; i < n; i++ {
		// 在子数组中间的时候
		if curVale+nums[i] > nums[i] {
			curVale = curVale + nums[i]
		} else {
			curVale = nums[i]
		}
		if curVale > resValue {
			resValue = curVale
		}

		// 从右向左检索时使用
		leftSum += nums[i]
		if leftSum > leftMax[i-1] {
			leftMax = append(leftMax, leftSum)
		} else {
			leftMax = append(leftMax, leftMax[i-1])
		}
	}

	rightSum := 0
	// 从后向前检索
	for j := n - 1; j > 0; j-- {
		rightSum += nums[j]
		if rightSum+leftMax[j-1] > resValue {
			resValue = rightSum + leftMax[j-1]
		}
	}

	return resValue
}

func Test918(t *testing.T) {
	//arr := []int{-2, 3, 2, -3, 3}
	arr1 := []int{1, -2, 3, 2}
	maxValue := maxSubarraySumCircular(arr1)
	fmt.Println(maxValue)
}
