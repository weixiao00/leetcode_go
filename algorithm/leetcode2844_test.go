package algorithm

import (
	"fmt"
	"testing"
)

// 贪心
// 被25整除。结尾是00，25，50，75
func minimumOperations(num string) int {

	if len(num) == 0 {
		return 0
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	zero := false
	five := false

	res := len(num)
	for i := len(num) - 1; i >= 0; i-- {
		if (num[i] == '0' || num[i] == '5') && zero {
			// 剩下的i+1+1
			res = getMin(res, len(num)-i-2)
		}
		if num[i] == '0' {
			zero = true
		}
		if (num[i] == '2' || num[i] == '7') && five {
			res = getMin(res, len(num)-i-2)
		}
		if num[i] == '5' {
			five = true
		}
	}

	if zero {
		res = getMin(res, len(num)-1)
	}

	return res
}

func Test2844(t *testing.T) {
	num := "2255047"
	res := minimumOperations(num)
	fmt.Println(res)
}
