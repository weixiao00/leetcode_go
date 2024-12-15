package algorithm

import (
	"fmt"
	"testing"
)

// 模拟
func semiOrderedPermutation(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := 0

	for i := range nums {
		if nums[i] == 1 {
			left = i
		}
		if nums[i] == len(nums) {
			right = i
		}
	}

	res := 0
	res += len(nums) - 1 - right
	res += left
	if right < left {
		res -= 1
	}

	return res
}

func Test2717(t *testing.T) {
	nums := []int{2, 5, 4, 1, 3}
	res := semiOrderedPermutation(nums)
	fmt.Println(res)
}
