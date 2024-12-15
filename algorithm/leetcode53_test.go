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

// 求连续子数组的最大和,并返回对应的下标
// * 请输出最大和及对应的开始、结束下标位置。
// * 例如输入的数组为{1, -2, 3, 10, -4, 7, 2, -5}，
// * 和最大的子数组为{3，10，-4，7，2}，
// * 因此输出为该数组的和18，开始、结束下标位置分别为2、6。
func maxSubArray2(nums []int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	res := math.MinInt
	curVal := 0
	beginIdx := 0
	endIdx := 0
	tmpBeginIdx := 0
	tmpEndIdx := 0
	for i, num := range nums {
		if curVal < 0 {
			curVal = 0
			tmpBeginIdx = i
			tmpEndIdx = i
		}
		curVal += num
		if curVal > res {
			res = curVal
			beginIdx = tmpBeginIdx
			endIdx = tmpEndIdx
		}
		tmpEndIdx++
	}

	return []int{res, beginIdx, endIdx}
}

func Test53(t *testing.T) {
	arr := []int{1, -2, 3, 10, -4, 7, 2, -5}
	maxValue := maxSubArray2(arr)
	fmt.Println(maxValue)
}
