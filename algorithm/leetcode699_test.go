package algorithm

import (
	"fmt"
	"testing"
)

// 暴力
func fallingSquares(positions [][]int) []int {

	if len(positions) == 0 {
		return []int{}
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	height := make([]int, len(positions))
	height[0] = positions[0][1]
	for i := 1; i < len(positions); i++ {
		left1 := positions[i][0]
		right1 := positions[i][0] + positions[i][1]
		height[i] = positions[i][1]
		for j := 0; j < i; j++ {
			left2 := positions[j][0]
			right2 := positions[j][0] + positions[j][1]
			// 有重叠
			if left1 < right2 && right1 > left2 {
				height[i] = getMax(height[i], height[j]+positions[i][1])
			}
		}
	}

	for i := 1; i < len(height); i++ {
		height[i] = getMax(height[i-1], height[i])
	}

	return height
}

func Test699(t *testing.T) {
	positions := [][]int{{1, 2}, {2, 3}, {6, 1}}
	res := fallingSquares(positions)
	fmt.Println(res)
}
