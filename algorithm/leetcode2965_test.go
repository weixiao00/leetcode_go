package algorithm

import (
	"fmt"
	"testing"
)

func findMissingAndRepeatedValues(grid [][]int) []int {

	res := make([]int, 2)
	if len(grid) == 0 {
		return res
	}

	row := len(grid)
	column := len(grid[0])
	counts := make([]int, row*column+1)
	for _, sg := range grid {
		for _, val := range sg {
			counts[val] += 1
		}
	}

	for i, val := range counts {
		if i == 0 {
			continue
		}
		if val == 2 {
			res[0] = i
		}
		if val == 0 {
			res[1] = i
		}
	}

	return res
}

func Test2965(t *testing.T) {
	grid := [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}
	res := findMissingAndRepeatedValues(grid)
	fmt.Println(res)
}
