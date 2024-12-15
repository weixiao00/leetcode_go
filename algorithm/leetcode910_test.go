package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func smallestRangeII(nums []int, k int) int {

	if len(nums) == 0 {
		return 0
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sort.Ints(nums)
	res := 0
	max := nums[len(nums)-1]
	min := nums[0]
	res = max - min

	for i := 0; i < len(nums)-1; i++ {
		a := nums[i]
		b := nums[i+1]
		res = getMin(getMax(max-k, a+k)-getMin(min+k, b-k), res)
	}

	return res
}

func Test910(t *testing.T) {
	nums := []int{1, 3, 6}
	k := 3
	res := smallestRangeII(nums, k)
	fmt.Println(res)
}
