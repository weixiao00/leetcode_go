package algorithm

import (
	"fmt"
	"testing"
)

// 贪心
func getSmallestString(s string, k int) string {

	if len(s) == 0 {
		return s
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	res := ""
	for i, ss := range s {
		if k == 0 {
			res += s[i:]
			return res
		}
		afterStep := int('z' - ss)
		beforeStep := int(ss - 'a')
		minStep := getMin(beforeStep, afterStep+1)
		// 变不到a，就往小的变
		if minStep > k {
			res += string(rune(int(ss) - k))
			k = 0
		} else {
			res += string('a')
			k -= minStep
		}
	}

	return res
}

func Test3106(t *testing.T) {
	s := "xaxcd"
	k := 4
	res := getSmallestString(s, k)
	fmt.Println(res)
}
