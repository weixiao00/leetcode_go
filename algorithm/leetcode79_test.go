package algorithm

import (
	"fmt"
	"testing"
)

// dfs, 大山的数量
func exist(board [][]byte, word string) bool {

	if len(board) == 0 {
		return false
	}

	row := len(board)
	column := len(board[0])

	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, column)
	}
	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		if index == len(word) {
			return true
		}
		if i < 0 || i >= row || j < 0 || j >= column || board[i][j] != word[index] {
			return false
		}
		if visited[i][j] {
			return false
		}
		res := false
		visited[i][j] = true
		res = res || dfs(i+1, j, index+1)
		res = res || dfs(i-1, j, index+1)
		res = res || dfs(i, j+1, index+1)
		res = res || dfs(i, j-1, index+1)
		visited[i][j] = false

		return res
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func Test79(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCB"
	res := exist(board, word)
	fmt.Println(res)
}
