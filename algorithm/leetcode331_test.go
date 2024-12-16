package algorithm

import (
	"fmt"
	"testing"
)

// 栈. 每一个节点看作一个曹位
func isValidSerialization(preorder string) bool {

	if len(preorder) == 0 {
		return true
	}

	stack := make([]int, 0)
	stack = append(stack, 1)

	i := 0
	for i < len(preorder) {
		if len(stack) == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
			continue
		} else if preorder[i] == '#' {
			stack[0]--
			if stack[0] == 0 {
				stack = stack[1:]
			}
			i++
		} else {

			//"9,#,92,#,#" 可能会有这种case
			for i < len(preorder) && preorder[i] != ',' {
				i++
			}

			stack[0]--
			if stack[0] == 0 {
				stack = stack[1:]
			}
			stack = append(stack, 2)
		}
	}

	return len(stack) == 0
}

func Test331(t *testing.T) {
	preorder := "9,#,92,#,#"
	res := isValidSerialization(preorder)
	fmt.Println(res)
}
