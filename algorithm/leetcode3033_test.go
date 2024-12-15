package algorithm

import (
	"fmt"
	"testing"
)

func modifiedMatrix(matrix [][]int) [][]int {

	if len(matrix) == 0 {
		return matrix
	}
	m := len(matrix)
	n := len(matrix[0])

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < n; i++ {
		max := 0
		arr := make([]int, 0)
		for j := 0; j < m; j++ {
			if matrix[j][i] == -1 {
				arr = append(arr, j)
			}
			max = getMax(max, matrix[j][i])
		}
		for _, k := range arr {
			matrix[k][i] = max
		}
	}

	return matrix
}

func Test3033(t *testing.T) {
	matrix := [][]int{{1, 2, -1}, {4, -1, 6}, {7, 8, 9}}
	matrix = modifiedMatrix(matrix)
	fmt.Println(matrix)
}
