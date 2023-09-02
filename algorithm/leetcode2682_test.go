package algorithm

import (
	"fmt"
	"testing"
)

func circularGameLosers(n int, k int) []int {
	visit := make([]bool, n)
	i := k
	j := 0
	for true {
		if visit[j] {
			break
		}
		visit[j] = true
		j = (i + j) % n
		i += k
	}

	res := make([]int, 0, n)
	for i, m := range visit {
		if !m {
			res = append(res, i+1)
		}
	}

	return res
}

func Test2682(t *testing.T) {
	res := circularGameLosers(5, 2)
	for _, i := range res {
		fmt.Println(i)
	}
}
