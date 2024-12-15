package algorithm

import (
	"fmt"
	"testing"
)

// 计算行列1的个数(row-1)*(col-1)
func numberOfRightTriangles(grid [][]int) int64 {

	if len(grid) == 0 {
		return 0
	}

	var res int64 = 0

	m := len(grid)
	n := len(grid[0])
	row := make([]int, m)
	col := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			row[i] += grid[i][j]
		}
	}

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			col[j] += grid[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res += int64((row[i] - 1) * (col[j] - 1))
			}
		}
	}

	return res
}

func Test3128(t *testing.T) {
	grid := [][]int{{0, 1, 0}, {0, 1, 1}, {0, 1, 0}}
	res := numberOfRightTriangles(grid)
	fmt.Println(res)
}
