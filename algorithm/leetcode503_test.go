package algorithm

import (
	"fmt"
	"testing"
)

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	if len(nums) == 0 {
		return res
	}
	for i := range res {
		res[i] = -1
	}

	n := len(nums)*2 - 1

	// 单调栈
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i%len(nums)] {
			res[stack[len(stack)-1]] = nums[i%len(nums)]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%len(nums))
	}

	return res
}

func Test503(t *testing.T) {
	nums := []int{1, 2, 3, 4, 3}
	res := nextGreaterElements(nums)
	fmt.Println(res)
}
