package algorithm

import (
	"fmt"
	"testing"
)

func maxOperations(nums []int) int {

	res := 0
	length := len(nums)
	if length == 0 {
		return 0
	}

	count := 0

	for len(nums) > 1 {
		i := nums[0] + nums[1]
		if count == 0 || count == i {
			count = i
			nums = nums[2:]
			res++
			continue
		}
		return res
	}

	return res
}

func Test3038(t *testing.T) {
	nums := []int{3, 2, 1, 4, 5}
	operations := maxOperations(nums)
	fmt.Println(operations)
}
