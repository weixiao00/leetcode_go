package algorithm

import (
	"fmt"
	"testing"
)

// 异或。技巧
func singleNumber(nums []int) int {

	res := 0
	for i := range nums {
		res ^= nums[i]
	}

	return res
}

func Test136(t *testing.T) {
	nums := []int{2, 2, 1}
	res := singleNumber(nums)
	fmt.Println(res)
}
