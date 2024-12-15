package algorithm

import (
	"fmt"
	"testing"
)

func containsDuplicate(nums []int) bool {

	if len(nums) < 2 {
		return false
	}
	mp := make(map[int]bool)
	for _, num := range nums {
		if mp[num] {
			return true
		}
		mp[num] = true
	}

	return false
}

func Test217(t *testing.T) {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	res := containsDuplicate(nums)
	fmt.Println(res)
}
