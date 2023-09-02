package algorithm

import (
	"fmt"
	"testing"
)

func countServers(grid [][]int) int {

	if grid == nil {
		return 0
	}
	rowSize := len(grid)
	columnSize := len(grid[0])
	rowMap := make(map[int]int)
	columnMap := make(map[int]int)

	for i := 0; i < rowSize; i++ {
		for j := 0; j < columnSize; j++ {
			if grid[i][j] == 1 {
				rowMap[i] = rowMap[i] + 1
				columnMap[j] = columnMap[j] + 1
			}
		}
	}

	res := 0
	for i := 0; i < rowSize; i++ {
		for j := 0; j < columnSize; j++ {
			if grid[i][j] == 1 && (rowMap[i] > 1 || columnMap[j] > 1) {
				res++
			}
		}
	}

	return res
}

func Test1267(t *testing.T) {
	arr := [][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}
	servers := countServers(arr)
	fmt.Println(servers)
}
