package algorithm

import (
	"fmt"
	"math"
	"testing"
)

// 可以用暴力，时间复杂度为O(n*n)
// 前缀和+hash 或者使用滑动窗口
// 等于target
func minSubArrayLen1(target int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	res := math.MaxInt32
	countIdxMap := make(map[int]int)
	countIdxMap[0] = -1

	count := 0
	for i, num := range nums {
		count += num
		preIdx, ok := countIdxMap[count-target]
		if ok {
			res = getMin(res, i-preIdx)
		}
		countIdxMap[count] = i
	}

	return res
}

// 大于等于target
// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	res := math.MaxInt32
	count := 0
	j := 0

	for i := 0; i < len(nums); i++ {
		count += nums[i]

		for j <= i && count >= target {
			res = getMin(res, i-j+1)
			count -= nums[j]
			j++
		}
	}
	if res == math.MaxInt32 {
		return 0
	}

	return res
}

func TestLcp008(t *testing.T) {
	target := 11
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}
