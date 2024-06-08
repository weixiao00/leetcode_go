package algorithm

import (
	"fmt"
	"testing"
)

func moveZeroes1(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	start := 0
	end := 0
	for i := range nums {
		if nums[start] != 0 {
			start++
			end++
		}
		if start >= len(nums) {
			break
		}
		if nums[start] == 0 && nums[i] != 0 {
			end = i
		}
		if start < end {
			exchangeArr(nums, start, end)
			i = start
			start = i
			end = i
		}
	}

	return nums
}

func moveZeroes(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	start := 0
	end := 0
	for i, val := range nums {
		if val != 0 {
			continue
		}
		start = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] != 0 {
				end = j
				break
			}
		}
		if start < end {
			exchangeArr(nums, start, end)
			i = end + 1
		}
	}

	return nums
}

func exchangeArr(arr []int, start, end int) {
	tmp := arr[start]
	arr[start] = arr[end]
	arr[end] = tmp
}

func Test123(t *testing.T) {
	arr := []int{45192, 0, -659, -52359, -99225, -75991, 0, -15155, 27382, 59818, 0, -30645, -17025, 81209, 887, 64648}
	ints := moveZeroes1(arr)
	fmt.Println(ints)
}
