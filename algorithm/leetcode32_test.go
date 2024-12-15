package algorithm

import (
	"fmt"
	"testing"
)

func longestValidParentheses(s string) int {

	if len(s) == 0 {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	stack := make([]int, 0)
	stack = append(stack, -1)

	res := 0
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[0 : len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = getMax(res, i-stack[len(stack)-1])
			}
		}
	}

	return res
}

func Test32(t *testing.T) {
	s := "())"
	res := longestValidParentheses(s)
	fmt.Println(res)
}
