package algorithm

import (
	"fmt"
	"testing"
)

// 分类讨论
// 抓到皇后最多两步。
// 如果象在车和皇后中间。移走象车就可以攻击皇后了
func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {

	var abs func(a int) int
	abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 1. 在同一行，象不在中间
	if a == e && (c != e || d <= getMin(b, f) || d >= getMax(b, f)) {
		return 1
	}

	// 2. 在同一列，象不在中间
	if b == f && (d != f || c <= getMin(a, e) || c >= getMax(a, e)) {
		return 1
	}

	// 3. 在同一对角线，象不在中间
	if abs(c-e) == abs(d-f) && ((a-e)*(c-e) != (b-f)*(d-f) || (a < getMin(c, e) || a > getMax(c, e))) {
		return 1
	}

	return 2
}

func Test3001(t *testing.T) {
	a := 4
	b := 6
	c := 1
	d := 1
	e := 5
	f := 5
	res := minMovesToCaptureTheQueen(a, b, c, d, e, f)
	fmt.Println(res)
}
