package algorithm

import (
	"fmt"
	"testing"
)

func maximumCount2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	left := 0
	right := length - 1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] < 0 {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	neg := left
	left = 0
	right = length - 1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] > 0 {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return getMax(neg, length-left)
}

func Test2529(t *testing.T) {
	arr := []int{-2, -1, -1, 1, 2, 3}
	fmt.Println(maximumCount2(arr))
}

func maximumCount(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	a := 0
	b := 0
	for _, i := range nums {
		if i < 0 {
			b++
		} else if i > 0 {
			a++
		}
	}

	return getMax(a, b)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
