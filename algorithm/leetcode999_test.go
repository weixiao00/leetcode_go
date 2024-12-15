package algorithm

import (
	"fmt"
	"testing"
)

// 模拟
// 最多可以吃四个
func numRookCaptures(board [][]byte) int {

	m := len(board)
	n := len(board[0])

	row := 0
	column := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'R' {
				row = i
				column = j
			}
		}
	}

	res := 0
	tmp := column
	for i := 1; i < n; i++ {
		if tmp-i >= 0 && board[row][tmp-i] == 'B' {
			break
		}
		if tmp-i >= 0 && board[row][tmp-i] == 'p' {
			res++
			break
		}
	}
	tmp = column
	for i := 1; i < n; i++ {
		if tmp+i < n && board[row][tmp+i] == 'B' {
			break
		}
		if tmp+i < n && board[row][tmp+i] == 'p' {
			res++
			break
		}
	}
	tmp = row
	for i := 1; i < m; i++ {
		if tmp+i < m && board[tmp+i][column] == 'B' {
			break
		}
		if tmp+i < m && board[tmp+i][column] == 'p' {
			res++
			break
		}
	}
	tmp = row
	for i := 1; i < m; i++ {
		if tmp-i >= 0 && board[tmp-i][column] == 'B' {
			break
		}
		if tmp-i >= 0 && board[tmp-i][column] == 'p' {
			res++
			break
		}
	}

	return res
}

func Test999(t *testing.T) {

	chessBoard := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	res := numRookCaptures(chessBoard)
	fmt.Println(res)
}
