package algorithm

import (
	"fmt"
	"testing"
)

// 前缀和+hash
// 可以不用nums[i]。使用+1，-1
func findMaxLength(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	mp := make(map[int]int)
	mp[0] = -1

	res := 0
	counter := 0
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}

		preIdx, ok := mp[counter]
		if ok {
			res = getMax(res, i-preIdx)
		} else {
			mp[counter] = i
		}
	}

	return res
}

func TestLcr011(t *testing.T) {
	nums := []int{0, 1, 0}
	res := findMaxLength(nums)
	fmt.Println(res)
}
