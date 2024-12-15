package algorithm

import (
	"fmt"
	"testing"
)

// 模拟+找规律
func rotate1(matrix [][]int) {

	if len(matrix) == 0 {
		return
	}

	m := len(matrix)
	n := len(matrix[0])

	// 辅助数组
	assist := make([][]int, m)
	for i := range assist {
		assist[i] = make([]int, n)
	}

	column := len(matrix)
	for i := 0; i < m; i++ {
		column -= 1
		idx := 0
		for j := 0; j < n; j++ {
			assist[idx][column] = matrix[i][j]
			idx++
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = assist[i][j]
		}
	}
}

func Test48(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate1(matrix)
	fmt.Println(matrix)
}
