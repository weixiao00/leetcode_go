package algorithm

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {

	if len(board) == 0 {
		return false
	}
	// 八个方向
	row := []int{1, 1, 0, -1, -1, -1, 0, 1}
	column := []int{0, 1, 1, 1, 0, -1, -1, -1}
	for i := 0; i < 8; i++ {
		if check1(board, rMove, cMove, color, row[i], column[i]) {
			return true
		}
	}
	return false
}

func check1(board [][]byte, rMove int, cMove int, color byte, row, column int) bool {
	step := 1
	rMove += row
	cMove += column
	for rMove >= 0 && rMove < 8 && cMove >= 0 && cMove < 8 {
		if board[rMove][cMove] == '.' {
			return false
		}
		if step == 1 {
			if board[rMove][cMove] == color {
				return false
			}
		} else {
			if board[rMove][cMove] == color {
				return true
			}
		}
		step++
		rMove += row
		cMove += column
	}

	return false
}
