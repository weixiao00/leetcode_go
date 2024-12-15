package algorithm

import (
	"fmt"
	"testing"
)

// mp的hash
func majorityElement(nums []int) int {

	mp := make(map[int]int)

	for i := range nums {
		mp[nums[i]]++
	}

	for key, val := range mp {
		if val > len(nums)/2 {
			return key
		}
	}

	return -1
}

func majorityElement1(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	candidate := -1
	count := 0

	for i := range nums {
		if count == 0 {
			candidate = nums[i]
		}
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func Test169(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	res := majorityElement1(nums)
	fmt.Println(res)
}
