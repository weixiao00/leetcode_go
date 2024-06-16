package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func maxScore(nums []int, x int) int64 {

	if len(nums) == 0 {
		return 0
	}
	// 用两个变量存储奇偶的最大值
	// 奇数
	oddNumber := math.MinInt32
	// 偶数
	evenNumber := math.MinInt32
	if nums[0]%2 == 0 {
		evenNumber = nums[0]
	} else {
		oddNumber = nums[0]
	}

	for i, num := range nums {
		if i == 0 {
			continue
		}
		// 说明是偶数
		if num%2 == 0 {
			evenNumber = getMax11(num+evenNumber, oddNumber+num-x)
		} else {
			oddNumber = getMax11(oddNumber+num, num+evenNumber-x)
		}
	}

	return int64(getMax11(oddNumber, evenNumber))
}

func getMax11(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test2786(t *testing.T) {
	nums := []int{1, 1000000}
	x := 1000000
	score := maxScore(nums, x)
	fmt.Println(score)
}
