package algorithm

import (
	"fmt"
	"testing"
)

// 标记数组
func setZeroes(matrix [][]int) {

	if len(matrix) == 0 {
		return
	}

	rowMap := make(map[int]bool)
	columnMap := make(map[int]bool)

	row := len(matrix)
	column := len(matrix[0])

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j] == 0 {
				rowMap[i] = true
				columnMap[j] = true
			}
		}
	}

	// 每行都变成0
	for i := 0; i < row; i++ {
		if rowMap[i] {
			for j := 0; j < column; j++ {
				matrix[i][j] = 0
			}
		}
	}

	// 每行都变成0
	for j := 0; j < column; j++ {
		if columnMap[j] {
			for i := 0; i < row; i++ {
				matrix[i][j] = 0
			}
		}
	}
}

func Test73(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
