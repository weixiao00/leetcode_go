package algorithm

import (
	"fmt"
	"testing"
)

func maxIncreaseKeepingSkyline(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	m := len(grid)
	n := len(grid[0])

	rowMax := make([]int, m)
	columnMax := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rowMax[i] = getMax(rowMax[i], grid[i][j])
			columnMax[j] = getMax(columnMax[j], grid[i][j])
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += getMin(rowMax[i], columnMax[j]) - grid[i][j]
		}
	}

	return res
}

func Test807(t *testing.T) {
	grid := [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	res := maxIncreaseKeepingSkyline(grid)
	fmt.Println(res)
}
