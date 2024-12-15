package algorithm

import (
	"fmt"
	"testing"
)

func maxArea1(height []int) int {

	if len(height) < 2 {
		return 0
	}

	res := 0
	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	left := 0
	right := len(height) - 1

	for left < right {
		if height[left] < height[right] {
			res = getMax(res, (right-left)*height[left])
			left++
		} else {
			res = getMax(res, (right-left)*height[right])
			right--
		}
	}

	return res
}

func Test11(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea1(height)
	fmt.Println(res)
}
