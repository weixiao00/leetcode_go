package algorithm

import (
	"fmt"
	"testing"
)

func losingPlayer(x int, y int) string {

	ten := y / 4

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	if getMin(x, ten)&1 == 0 {
		return "Bob"
	}
	return "Alice"
}

func Test3222(t *testing.T) {
	x := 4
	y := 11
	res := losingPlayer(x, y)
	fmt.Println(res)
}
