package algorithm

func diagonalSum(mat [][]int) int {
	res := 0
	if mat == nil {
		return res
	}

	row := len(mat)
	column := len(mat[0])

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if i == j || i+j == row-1 {
				res += mat[i][j]
			}
		}
	}

	return res
}
