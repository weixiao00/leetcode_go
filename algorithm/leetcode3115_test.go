package algorithm

import (
	"fmt"
	"testing"
)

func maximumPrimeDifference(nums []int) int {

	if len(nums) < 2 {
		return 0
	}

	var prime func(num int) bool
	prime = func(num int) bool {
		if num == 1 {
			return false
		}
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	first := -1
	last := -1
	for i, num := range nums {
		if !prime(num) {
			continue
		}
		if first == -1 {
			first = i
		}
		last = i
	}

	return last - first
}

func Test3115(t *testing.T) {
	nums := []int{4, 2, 9, 5, 3}
	res := maximumPrimeDifference(nums)
	fmt.Println(res)
}
