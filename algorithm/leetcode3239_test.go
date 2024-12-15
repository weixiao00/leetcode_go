package algorithm

import (
	"fmt"
	"testing"
)

// 双指针
func minFlips(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	row := len(grid)
	column := len(grid[0])

	rowRes := 0
	for i := 0; i < row; i++ {
		left := 0
		right := column - 1
		for left < right {
			if grid[i][left] != grid[i][right] {
				rowRes++
			}
			left++
			right--
		}
	}

	columnRes := 0
	for i := 0; i < column; i++ {
		top := 0
		bottom := row - 1
		for top < bottom {
			if grid[top][i] != grid[bottom][i] {
				columnRes++
			}
			top++
			bottom--
		}

	}

	if rowRes > columnRes {
		return columnRes
	}
	return rowRes
}

func Test3239(t *testing.T) {
	grid := [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	res := minFlips(grid)
	fmt.Println(res)
}
