package algorithm

import (
	"fmt"
	"testing"
)

// 二分查找
func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	n := len(matrix[0])

	i := 0
	j := n - 1
	for i < m && j >= 0 {
		if target == matrix[i][j] {
			return true
		} else if target < matrix[i][j] {
			j--
		} else {
			i++
		}
	}

	return false
}

func Test74(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	res := searchMatrix(matrix, target)
	fmt.Println(res)
}
