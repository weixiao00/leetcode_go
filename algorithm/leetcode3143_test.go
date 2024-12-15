package algorithm

import (
	"fmt"
	"testing"
)

// 计算次小的距离
func maxPointsInsideSquare(points [][]int, s string) int {

	if len(points) == 0 || len(s) == 0 {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var abs func(a int) int
	abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	min1 := make([]int, 26)
	for i := range min1 {
		min1[i] = 10000000001
	}
	min2 := 10000000001

	n := len(points)
	for i := 0; i < n; i++ {
		x := points[i][0]
		y := points[i][1]
		j := s[i] - 'a'
		d := getMax(abs(x), abs(y))
		if d < min1[j] {
			min2 = getMin(min2, min1[j])
			min1[j] = d
		} else if d >= min1[j] && d < min2 {
			min2 = d
		}
	}

	res := 0
	for i := range min1 {
		if min1[i] < min2 {
			res++
		}
	}

	return res
}

func Test3143(t *testing.T) {
	x := 'b'
	fmt.Println(x - 'a')
}
