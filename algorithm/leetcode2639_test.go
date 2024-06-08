package algorithm

import (
	"fmt"
	"testing"
)

func findColumnWidth(grid [][]int) []int {
	if len(grid) == 0 {
		return []int{}
	}

	// 获取长度
	var getLen func(val int) int
	getLen = func(val int) int {
		len := 0
		if val <= 0 {
			len++
		}
		for val != 0 {
			val /= 10
			len++
		}
		return len
	}

	// 获取最大值
	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	row := len(grid)
	column := len(grid[0])
	res := make([]int, 0, column)
	for j := 0; j < column; j++ {
		len := 0
		for i := 0; i < row; i++ {
			val := grid[i][j]
			len = getMax(len, getLen(val))
		}
		res = append(res, len)
	}
	return res
}

func Test2639(t *testing.T) {
	grid := [][]int{{0}}
	width := findColumnWidth(grid)
	fmt.Println(width)
}
