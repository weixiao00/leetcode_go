package algorithm

import (
	"fmt"
	"testing"
)

func findLUSlength1(strs []string) int {
	if len(strs) == 0 {
		return -1
	}

	var isSub func(str1, str2 string) bool
	isSub = func(str1, str2 string) bool {
		// 这样的话子字符串得是连续的
		//s1 := strings.Contains(str1, str2)
		//s2 := strings.Contains(str2, str1)
		//return s1 || s2
		if str1 == str2 {
			return true
		}
		idx1 := 0
		idx2 := 0
		for idx1 < len(str1) && idx2 < len(str2) {
			if str1[idx1] == str2[idx2] {
				idx1++
			}
			idx2++
		}

		return len(str1) == idx1
	}

	res := -1
	for i := range strs {
		checked := true
		for j := range strs {
			if i != j && isSub(strs[i], strs[j]) {
				checked = false
				break
			}
		}

		if checked {
			res = getMax15(res, len(strs[i]))
		}
	}

	return res
}

func getMax15(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test522(t *testing.T) {
	strs := []string{"aabbcc", "aabbcc", "cb", "abc"}
	slength1 := findLUSlength1(strs)
	fmt.Println(slength1)
}
