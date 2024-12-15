package algorithm

import (
	"fmt"
	"testing"
)

// 全排列
func combinationSum(candidates []int, target int) [][]int {

	if len(candidates) == 0 {
		return [][]int{}
	}

	res := make([][]int, 0)
	var dfs func(count []int, sum, k int)
	dfs = func(count []int, sum, k int) {
		if sum > target {
			return
		}

		if sum == target {
			tmp := make([]int, len(count))
			copy(tmp, count)
			res = append(res, tmp)
			return
		}

		for i, num := range candidates {
			if i < k {
				continue
			}
			sum += num
			count = append(count, num)
			dfs(count, sum, i)
			sum -= num
			count = count[0 : len(count)-1]
		}
	}

	dfs([]int{}, 0, 0)

	return res
}

func Test39(t *testing.T) {
	candidates := []int{2, 3, 5}
	target := 8
	res := combinationSum(candidates, target)
	fmt.Println(res)
}
