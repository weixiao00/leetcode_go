package algorithm

import (
	"fmt"
	"testing"
)

// 暴力O(N*N)
// 单调栈O(N)
func largestRectangleArea(heights []int) int {

	if len(heights) == 0 {
		return 0
	}

	n := len(heights)

	// 每个下表左边和右边第一个小于的
	left := make([]int, n)
	right := make([]int, n)

	stack := make([]int, 0)
	for i := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	for i := range heights {
		res = getMax(res, (right[i]-left[i]-1)*heights[i])
	}

	return res
}

// 单次遍历
func largestRectangleArea1(heights []int) int {

	if len(heights) == 0 {
		return 0
	}

	n := len(heights)

	// 每个下表左边和右边第一个小于的
	left := make([]int, n)
	right := make([]int, n)
	for i := range right {
		right[i] = n
	}

	stack := make([]int, 0)
	for i := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	for i := range heights {
		res = getMax(res, (right[i]-left[i]-1)*heights[i])
	}

	return res
}

func Test84(t *testing.T) {
	heights := []int{2, 4}
	res := largestRectangleArea1(heights)
	fmt.Println(res)
}
