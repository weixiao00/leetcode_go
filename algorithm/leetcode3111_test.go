package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

// 贪心
func minRectanglesToCoverPoints(points [][]int, w int) int {

	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	res := 0
	wRightIdx := -1
	for _, point := range points {
		if wRightIdx == -1 || point[0] > wRightIdx {
			wRightIdx = point[0] + w
			res++
		}
	}

	return res
}

func Test3111(t *testing.T) {
	points := [][]int{{2, 1}, {1, 0}, {1, 4}, {1, 8}, {3, 5}, {4, 6}}
	w := 1
	res := minRectanglesToCoverPoints(points, w)
	fmt.Println(res)
}
