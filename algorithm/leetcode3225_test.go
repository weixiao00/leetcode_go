package algorithm

import (
	"fmt"
	"testing"
)

// 数学题
// 没做出来，dfs挺难
func canReachCorner(xCorner int, yCorner int, circles [][]int) bool {

	yTop := true
	yBottom := true
	xLeft := true
	xRight := true

	for _, circle := range circles {

		x := circle[0]
		y := circle[1]
		r := circle[2]

		if x-r >= xCorner || x+r <= 0 || y-r >= yCorner || y+r <= 0 {
			continue
		}

		if x+r >= xCorner && x-r < xCorner {
			xRight = false
		}
		if x+r > 0 && x-r <= 0 {
			xLeft = false
		}
		if y+r >= yCorner && y-r < yCorner {
			yTop = false
		}
		if y+r > 0 && y-r <= 0 {
			yBottom = false
		}
	}

	if xLeft && yTop {
		return true
	}
	if xRight && yBottom {
		return true
	}

	return false
}

func Test3225(t *testing.T) {
	// 从中间走的
	X := 15
	Y := 15
	//[[2,1000,997],[1000,2,997]]
	circles := [][]int{{1, 99, 85}, {99, 1, 85}}
	res := canReachCorner(X, Y, circles)
	fmt.Println(res)
}
