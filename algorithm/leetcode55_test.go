package algorithm

import (
	"fmt"
	"testing"
)

// 贪心，就是往远跳
func canJump(nums []int) bool {

	if len(nums) < 2 {
		return true
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxIdx := 0
	for i, num := range nums {
		// 可以跳到这个坐标
		if i <= maxIdx {
			maxIdx = getMax(maxIdx, i+num)
		}
		// 大于说明跳不到这里。可以没有循环结束也是返回false
		//if i > maxIdx {
		//	return false
		//}
		if maxIdx >= len(nums)-1 {
			return true
		}
	}

	return false
}

// 自己的解法
// 没通过，可能会有多个0
func canJump1(nums []int) bool {

	if len(nums) < 2 {
		return true
	}

	zeroIdx := -1
	for i, num := range nums {
		if num == 0 {
			zeroIdx = i
		}
	}
	// 没有0
	if zeroIdx == -1 {
		return true
	}

	for i := 0; i < zeroIdx; i++ {
		if i+nums[i] > zeroIdx {
			return true
		}
	}
	return false
}

func Test55(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	res := canJump1(nums)
	fmt.Println(res)
}
