package algorithm

import (
	"fmt"
	"math"
	"testing"
)

// 暴力解法
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	MOD := 10000000007

	if len(nums) == 1 {
		return nums[0]
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		count := 1
		for j := i; j < len(nums); j++ {
			count *= nums[j]
			count %= MOD
			res = getMax(res, count)
		}
	}

	return res
}

// 动态规划或者叫贪心？
func maxProduct1(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	max := math.MinInt32
	min := 1
	res := 0
	for _, num := range nums {
		if num < 0 {
			tmp := max
			max = min
			min = tmp
		}
		max = getMax(num*max, num)
		min = getMin(num*min, num)
		res = getMax(res, max)
	}

	return res
}

// 动态规划，官方解法
func maxProduct2(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	maxDp := make([]int, len(nums))
	minDp := make([]int, len(nums))
	for i := range nums {
		maxDp[i] = nums[i]
		minDp[i] = nums[i]
	}

	res := maxDp[0]

	for i := 1; i < len(nums); i++ {
		maxDp[i] = getMax(nums[i], getMax(nums[i]*maxDp[i-1], nums[i]*minDp[i-1]))
		minDp[i] = getMin(nums[i], getMin(nums[i]*minDp[i-1], nums[i]*maxDp[i-1]))
		res = getMax(res, maxDp[i])
	}

	return res
}

func Test152(t *testing.T) {
	//1000000000
	nums := []int{3, -1, 4}
	res := maxProduct2(nums)
	fmt.Println(res)
}
