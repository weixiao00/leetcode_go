package algorithm

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func findValueOfPartition(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	res := math.MaxInt

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < len(nums)-1; i++ {
		res = getMin(res, nums[i+1]-nums[i])
	}

	return res
}

func Test2740(t *testing.T) {

	nums := []int{100, 1, 10}
	res := findValueOfPartition(nums)
	fmt.Println(res)
}
