package algorithm

import (
	"fmt"
	"testing"
)

func countTriplets(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	count := make([]int, 1<<16)

	for _, num1 := range nums {
		for _, num2 := range nums {
			count[num1&num2]++
		}
	}

	res := 0
	for _, num3 := range nums {
		for i := 0; i < 1<<16; i++ {
			if num3&i == 0 {
				res += count[i]
			}
		}
	}

	return res
}

func Test982(t *testing.T) {
	nums := []int{1, 1}
	res := countTriplets(nums)
	fmt.Println(res)
}
