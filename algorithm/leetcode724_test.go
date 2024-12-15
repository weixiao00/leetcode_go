package algorithm

import (
	"fmt"
	"testing"
)

func pivotIndex(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	// 总和
	totalNum := 0
	for _, num := range nums {
		totalNum += num
	}

	prefixNum := 0
	for i, num := range nums {
		if totalNum == 2*prefixNum+num {
			return i
		}
		prefixNum += num
	}

	return -1
}

func Test724(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	res := pivotIndex(nums)
	fmt.Println(res)
}
