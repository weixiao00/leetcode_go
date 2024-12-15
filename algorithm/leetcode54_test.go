package algorithm

import (
	"fmt"
	"testing"
)

// 模拟+辅助矩阵
func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return []int{}
	}

	m := len(matrix)
	n := len(matrix[0])

	// 辅助矩阵
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	res := make([]int, 0)
	steps := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	idx := 0
	i := 0
	j := 0
	// 循环m*n次
	// for i := 0; i < total; i++ {}
	for len(res) < m*n {
		res = append(res, matrix[i][j])
		visited[i][j] = true
		step := steps[idx%4]
		if i+step[0] == m || i+step[0] < 0 || j+step[1] == n ||
			j+step[1] < 0 || visited[i+step[0]][j+step[1]] {
			idx++
		}
		step = steps[idx%4]
		i += step[0]
		j += step[1]
	}

	return res
}

func Test54(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := spiralOrder(matrix)
	fmt.Println(res)
}
