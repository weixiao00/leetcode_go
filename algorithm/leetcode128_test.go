package algorithm

import (
	"fmt"
	"testing"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numMap := make(map[int]bool, 0)
	for _, num := range nums {
		numMap[num] = true
	}

	res := 0
	cur := 0
	for key := range numMap {
		if !numMap[key-1] {
			cur = 1

			for numMap[key+1] {
				cur++
				key++
			}
			res = getMax12(res, cur)
		}
	}

	return res
}

func getMax12(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test128(t *testing.T) {
	arr := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	consecutive := longestConsecutive(arr)
	fmt.Println(consecutive)
}
