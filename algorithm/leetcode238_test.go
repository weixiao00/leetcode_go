package algorithm

import (
	"fmt"
	"testing"
)

// 前缀积和后缀积
// 空间复杂度O(n*n) 可优化成O(n),O(1)
func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	leftProduct := make([]int, 0)
	rightProduct := make([]int, 0)

	left := 1
	for i := 0; i < len(nums); i++ {
		leftProduct = append(leftProduct, left)
		left *= nums[i]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		rightProduct = append(rightProduct, right)
		right *= nums[i]
	}

	answer := make([]int, 0)
	for i := range nums {
		answer = append(answer, leftProduct[i]*rightProduct[len(nums)-i-1])
	}

	return answer
}

// 空间复杂度O(1)
func productExceptSelf1(nums []int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	answer := make([]int, len(nums))
	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] = right * answer[i]
		right *= nums[i]
	}

	return answer
}

func Test238(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	res := productExceptSelf1(nums)
	fmt.Println(res)
}
