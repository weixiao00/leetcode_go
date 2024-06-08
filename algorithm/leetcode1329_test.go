package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func diagonalSort(mat [][]int) [][]int {
	if len(mat) == 0 {
		return mat
	}
	row := len(mat)
	column := len(mat[0])
	i := 0
	j := 0

	for i < row {
		innerRow := i
		innerColumn := 0
		arr := make([]int, 0)
		for innerRow < row && innerColumn < column {
			arr = append(arr, mat[innerRow][innerColumn])
			innerRow++
			innerColumn++
		}
		sort.Ints(arr)
		innerRow = i
		innerColumn = 0
		index := 0
		for innerRow < row && innerColumn < column {
			mat[innerRow][innerColumn] = arr[index]
			innerRow++
			innerColumn++
			index++
		}
		i++
	}

	for j < column {
		innerRow := 0
		innerColumn := j
		arr := make([]int, 0)
		for innerRow < row && innerColumn < column {
			arr = append(arr, mat[innerRow][innerColumn])
			innerRow++
			innerColumn++
		}
		sort.Ints(arr)
		innerRow = 0
		innerColumn = j
		index := 0
		for innerRow < row && innerColumn < column {
			mat[innerRow][innerColumn] = arr[index]
			innerRow++
			innerColumn++
			index++
		}
		j++
	}

	return mat
}

func Test1329(t *testing.T) {
	arr := [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}
	sari := diagonalSort(arr)
	fmt.Println(sari)
}
