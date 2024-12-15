package algorithm

import (
	"fmt"
	"testing"
)

// 模拟。找规律
func generate(numRows int) [][]int {

	res := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		row := make([]int, 0)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				up := res[i-1]
				row = append(row, up[j]+up[j-1])
			}
		}

		res = append(res, row)
	}

	return res
}

func Test118(t *testing.T) {
	numRows := 5
	res := generate(numRows)
	fmt.Println(res)
}
