package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func smallestRangeI(nums []int, k int) int {

	if len(nums) <= 1 {
		return 0
	}

	sort.Ints(nums)
	small := nums[0]
	big := nums[len(nums)-1]

	res := 0
	if big-small <= 2*k {
		res = 0
	} else {
		res = big - small - 2*k
	}

	return res
}

func Test908(t *testing.T) {
	nums := []int{2, 7, 2}
	res := smallestRangeI(nums, 1)
	fmt.Println(res)
}
