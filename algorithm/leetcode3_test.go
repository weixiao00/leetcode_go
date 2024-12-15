package algorithm

import (
	"fmt"
	"strings"
	"testing"
)

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	res := 0
	left := 0
	curVal := ""
	for right, ss := range s {
		for strings.Contains(curVal, string(ss)) {
			left++
			curVal = curVal[1:]
		}
		curVal += string(ss)
		res = getMax21(res, right-left+1)
	}

	return res
}

func getMax21(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test3(t *testing.T) {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
