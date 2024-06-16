package algorithm

import (
	"fmt"
	"testing"
)

func countBattleships(board [][]byte) int {

	res := 0
	if len(board) == 0 {
		return res
	}

	row := len(board)
	column := len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if board[i][j] == 'X' {
				if i > 0 && board[i-1][j] == 'X' {
					continue
				}
				if j > 0 && board[i][j-1] == 'X' {
					continue
				}
				res++
			}
		}
	}

	return res
}

func Test419(t *testing.T) {
	board := [][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}
	battleships := countBattleships(board)
	fmt.Println(battleships)
}
