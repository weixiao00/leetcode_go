package algorithm

import (
	"fmt"
	"testing"
)

// 岛屿的周长
func islandPerimeter(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	var dfs func(grid [][]int, i, j int) int
	dfs = func(grid [][]int, i, j int) int {
		if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 || grid[i][j] == 0 {
			return 1
		}

		if grid[i][j] == 2 {
			return 0
		}

		grid[i][j] = 2

		res := 0
		res += dfs(grid, i-1, j)
		res += dfs(grid, i+1, j)
		res += dfs(grid, i, j-1)
		res += dfs(grid, i, j+1)

		return res
	}

	res := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// 所有临近0的和4边的都算周长
				return dfs(grid, i, j)
			}
		}
	}

	return res
}

func Test463(t *testing.T) {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	res := islandPerimeter(grid)
	fmt.Println(res)
}
