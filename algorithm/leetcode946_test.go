package algorithm

import (
	"fmt"
	"testing"
)

func validateStackSequences(pushed []int, popped []int) bool {

	if len(pushed) == 0 {
		return true
	}

	i := 0
	j := 0
	stack := make([]int, 0, len(pushed))
	for j < len(popped) {
		// 一直出栈
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			j++
			stack = stack[0 : len(stack)-1]
		}

		// 可以入栈
		if i < len(pushed) {
			stack = append(stack, pushed[i])
			i++
		}

		// 都入了。出不来了
		if i == len(pushed) && len(stack) > 0 && stack[len(stack)-1] != popped[j] {
			return false
		}
	}

	return true
}

// 官方解法
func validateStackSequences1(pushed []int, popped []int) bool {

	n := len(pushed)

	stack := make([]int, 0)
	j := 0
	for i := 0; i < n; i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[0 : len(stack)-1]
			j++
		}
	}

	return len(stack) == 0
}

func Test946(t *testing.T) {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 3, 5, 1, 2}
	res := validateStackSequences(pushed, popped)
	fmt.Println(res)
}
