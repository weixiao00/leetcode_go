package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	res := math.MinInt
	curVal := 0
	for _, num := range nums {
		if curVal < 0 {
			curVal = 0
		}
		curVal += num
		if curVal > res {
			res = curVal
		}
	}
	return res
}

func Test53(t *testing.T) {
	arr := []int{5, 4, -1, 7, 8}
	maxValue := maxSubArray(arr)
	fmt.Println(maxValue)
}
