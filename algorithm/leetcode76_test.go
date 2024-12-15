package algorithm

import (
	"fmt"
	"math"
	"testing"
)

// 滑动窗口+比较个数
// 没有顺序，只比较个数就行
func minWindow(s string, t string) string {

	oriMap := make(map[rune]int)
	countMap := make(map[rune]int)

	for _, r := range []rune(t) {
		oriMap[r]++
	}

	var check func() bool
	check = func() bool {
		for key, val := range oriMap {
			if val > countMap[key] {
				return false
			}
		}
		return true
	}

	left := 0
	right := 0
	res := ""
	runes := []rune(s)
	aLen := math.MaxInt32
	for i := range runes {
		countMap[runes[i]]++
		for check() && left <= right {
			if aLen > right-left+1 {
				res = s[left : right+1]
				aLen = len(res)
			}
			countMap[runes[left]]--
			left++
		}
		right++
	}

	return res
}

func Test76(t *testing.T) {
	s := "ADOBECODEBANC"
	tt := "ABC"
	res := minWindow(s, tt)
	fmt.Println(res)
}
