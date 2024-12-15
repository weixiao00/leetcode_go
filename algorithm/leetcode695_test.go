package algorithm

import (
	"fmt"
	"testing"
)

// 岛屿的最大面积
func maxAreaOfIsland(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	var dfs func(grid [][]int, i, j int) int
	dfs = func(grid [][]int, i, j int) int {
		if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 || grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = 2

		res := 1
		res += dfs(grid, i-1, j)
		res += dfs(grid, i+1, j)
		res += dfs(grid, i, j-1)
		res += dfs(grid, i, j+1)

		return res
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// dfs将相邻的岛屿从1变成2.并且计算最大值
				res = getMax(res, dfs(grid, i, j))
			}
		}
	}

	return res
}

func Test695(t *testing.T) {
	grid := [][]int{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	res := maxAreaOfIsland(grid)
	fmt.Println(res)
}
