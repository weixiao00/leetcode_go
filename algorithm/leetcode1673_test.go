package algorithm

import (
	"container/list"
	"fmt"
	"testing"
)

// 单调栈.官方解法
func mostCompetitive(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	stack := list.New()

	len := len(nums)

	for i, num := range nums {
		for stack.Len() != 0 && len-i+stack.Len() > k && num < stack.Back().Value.(int) {
			// 移除栈顶元素
			stack.Remove(stack.Back())
		}
		stack.PushBack(num)
	}

	for stack.Len() > k {
		stack.Remove(stack.Back())
	}

	res := make([]int, 0)
	for stack.Len() > 0 {
		res = append(res, stack.Remove(stack.Front()).(int))
	}

	return res
}

func mostCompetitive1(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	res := make([]int, 0)
	if k == 0 {
		return res
	}

	return res
}

func Test1673(t *testing.T) {
	nums := []int{2, 4, 3, 3, 5, 4, 9, 6}
	k := 0
	competitive := mostCompetitive(nums, k)
	fmt.Println(competitive)
}
