package algorithm

import (
	"fmt"
	"testing"
)

// 计算每个点的入度和出度
func findJudge(n int, trust [][]int) int {

	input := make([]int, n+1)
	output := make([]int, n+1)
	for _, t := range trust {
		x := t[0]
		y := t[1]
		input[y]++
		output[x]++
	}
	for i := 1; i <= n; i++ {
		if input[i] == n-1 && output[i] == 0 {
			return i
		}
	}

	return -1
}

func Test997(t *testing.T) {

	n := 3
	trust := [][]int{{1, 3}, {2, 3}}
	res := findJudge(n, trust)
	fmt.Println(res)
}
