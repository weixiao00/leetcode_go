package algorithm

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// 官网解法，寻找下一个排列
func minimumMoves(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	more := make([][]int, 0)
	less := make([][]int, 0)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i][j] == 0 {
				less = append(less, []int{i, j})
			} else if grid[i][j] > 1 {
				for k := 2; k <= grid[i][j]; k++ {
					more = append(more, []int{i, j})
				}
			}
		}
	}

	var abs func(a, b int) int
	abs = func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// 因为默认排列是从小到大的
	var nextPermutation func(more [][]int) bool
	nextPermutation = func(more [][]int) bool {

		var isLess func(pairs1, pairs2 []int) bool
		isLess = func(pairs1, pairs2 []int) bool {
			return pairs1[0] < pairs2[0] || pairs1[0] == pairs2[0] && pairs1[1] < pairs2[1]
		}

		p := -1
		for i := 0; i < len(more)-1; i++ {
			if isLess(more[i], more[i+1]) {
				p = i
			}
		}

		// 没有下一个排列了
		if p == -1 {
			return false
		}

		q := -1
		// 从p+1开始获取大于nums[p] 最近的一个元素索引
		for j := p + 1; j < len(more); j++ {
			if isLess(more[p], more[j]) {
				q = j
			}
		}

		// 交换
		tmp := more[p]
		more[p] = more[q]
		more[q] = tmp

		// 交换排序
		i := p + 1
		j := len(more) - 1
		for i < j {
			tmp = more[i]
			more[i] = more[j]
			more[j] = tmp
			i++
			j--
		}
		return true
	}

	res := 0
	for i := range more {
		res += abs(more[i][0], less[i][0]) + abs(more[i][1], less[i][1])
	}
	for nextPermutation(more) {
		tmp := 0
		for i := range more {
			tmp += abs(more[i][0], less[i][0]) + abs(more[i][1], less[i][1])
		}
		res = getMin(res, tmp)
	}

	return res
}

// dfs，寻找所以不重复的排列
// 通过了
func minimumMoves1(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	more := make([][]int, 0)
	less := make([][]int, 0)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i][j] == 0 {
				less = append(less, []int{i, j})
			} else if grid[i][j] > 1 {
				for k := 2; k <= grid[i][j]; k++ {
					more = append(more, []int{i, j})
				}
			}
		}
	}

	var abs func(a, b int) int
	abs = func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	moreVisited := make([]bool, len(more))
	res := math.MaxInt32
	// 全排列
	var dfs func(curRes [][]int, idx int)
	dfs = func(curRes [][]int, idx int) {
		if idx == len(more) {
			tmp := 0
			for i := range curRes {
				tmp += abs(curRes[i][0], less[i][0]) + abs(curRes[i][1], less[i][1])
			}
			res = getMin(res, tmp)
			return
		}
		for m := 0; m < len(more); m++ {
			if moreVisited[m] || (m > 0 && reflect.DeepEqual(more[m], more[m-1]) && !moreVisited[m-1]) {
				continue
			}
			moreVisited[m] = true
			dfs(append(curRes, more[m]), idx+1)
			moreVisited[m] = false
		}
	}

	dfs([][]int{}, 0)

	return res
}

func Test2850(t *testing.T) {
	grid := [][]int{{1, 2, 2}, {1, 1, 0}, {0, 1, 1}}
	res := minimumMoves(grid)
	fmt.Println(res)
}
