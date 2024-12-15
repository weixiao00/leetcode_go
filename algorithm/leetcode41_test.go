package algorithm

import (
	"fmt"
	"testing"
)

// 原地处理
// 返回的结果要么是1到n。要么是n+1
func firstMissingPositive(nums []int) int {

	n := len(nums)

	var abs func(a int) int
	abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	// 回到本该属于它的位置
	for i := range nums {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}

	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}

// 交换排序。类似
func firstMissingPositive1(nums []int) int {

	n := len(nums)

	for i := range nums {
		// 相等就不交换了
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		}
	}

	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

func Test41(t *testing.T) {
	nums := []int{1, 1}
	res := firstMissingPositive(nums)
	fmt.Println(res)
}
