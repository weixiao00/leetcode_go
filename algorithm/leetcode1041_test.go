package algorithm

import (
	"fmt"
	"testing"
)

func Test1041(t *testing.T) {
	bounded := isRobotBounded("LLGGRR")
	fmt.Println(bounded)
}

func isRobotBounded(instructions string) bool {
	direct := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	directIndex := 0
	n := len(instructions)
	x := 0
	y := 0
	for i := 0; i < n; i++ {
		u := instructions[i]
		if u == 'G' {
			x += direct[directIndex][0]
			y += direct[directIndex][1]
		} else if u == 'L' {
			directIndex += 3
			directIndex %= 4
		} else {
			directIndex += 1
			directIndex %= 4
		}
	}

	// 回到原点||头不朝北
	return (x == 0 && y == 0) || directIndex != 0
}
