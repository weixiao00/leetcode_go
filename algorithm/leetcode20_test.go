package algorithm

import (
	"fmt"
	"testing"
)

// 栈
func isValid(s string) bool {

	stack := make([]rune, 0)

	runes := []rune(s)
	for _, r := range runes {
		if len(stack) == 0 || r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
			continue
		}

		if r == ')' && stack[len(stack)-1] != '(' {
			return false
		}
		if r == ']' && stack[len(stack)-1] != '[' {
			return false
		}
		if r == '}' && stack[len(stack)-1] != '{' {
			return false
		}
		stack = stack[0 : len(stack)-1]
	}

	return len(stack) == 0
}

func Test20(t *testing.T) {
	s := "()[]{}"
	res := isValid(s)
	fmt.Println(res)
}
