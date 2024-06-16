package algorithm

import (
	"fmt"
	"testing"
)

func generateParenthesis(n int) []string {
	var res []string
	if n <= 0 {
		return res
	}

	var dfs func(left, right int, s string)

	dfs = func(left, right int, s string) {
		if n*2 == len(s) {
			res = append(res, s)
		}
		if left > right {
			dfs(left, right+1, s+")")
		}
		if left < n {
			dfs(left+1, right, s+"(")
		}
	}

	dfs(0, 0, "")

	return res
}

func TestLcr085(t *testing.T) {
	parenthesis := generateParenthesis(3)
	fmt.Println(parenthesis)
}
