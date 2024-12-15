package algorithm

import (
	"fmt"
	"testing"
)

func maxSubArray_rep(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	res := nums[0]
	curVal := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		// 判断一下是否带着之前的值
		// 意思就是curVal是否大于0
		if curVal+num > num {
			curVal = curVal + num
		} else {
			curVal = num
		}
		if curVal > res {
			res = curVal
		}
	}
	return res
}

func Test53_rep(t *testing.T) {
	arr := []int{5, 4, -1, 7, 8}
	maxValue := maxSubArray_rep(arr)
	fmt.Println(maxValue)
}
