package algorithm

import (
	"fmt"
	"testing"
)

// 岛屿问题
// dfs搜索，类似于四叉树
func numIslands(grid [][]byte) int {

	if len(grid) == 0 {
		return 0
	}

	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 || grid[i][j] != '1' {
			return
		}
		grid[i][j] = 2
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}

	res := 0
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				// dfs将相邻的岛屿从1变成2
				dfs(grid, i, j)
				res++
			}
		}
	}

	return res
}

func Test200(t *testing.T) {
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	islands := numIslands(grid)
	fmt.Println(islands)
}
