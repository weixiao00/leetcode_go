package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func reverse(x int) int {

	var res int
	for x != 0 {
		val := x % 10
		x /= 10
		res = res*10 + val
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
	}

	return res
}

func Test7(t *testing.T) {
	x := -12
	res := reverse(x)
	fmt.Println(res)

	fmt.Println(x % 10)
}
