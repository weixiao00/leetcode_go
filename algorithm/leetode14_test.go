package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	res := ""
	min := math.MaxInt32
	for _, str := range strs {
		min = getMin(min, len(str))
	}

	for i := 0; i < min; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j-1][i] != strs[j][i] {
				return res
			}
		}
		res += string(strs[0][i])
	}

	return res
}

func Test14(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}
