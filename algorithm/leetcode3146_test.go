package algorithm

import (
	"fmt"
	"testing"
)

func findPermutationDifference(s string, t string) int {

	charIndexMap := make(map[rune]int)

	for i := range t {
		charIndexMap[rune(t[i])] = i
	}

	var abs func(a int) int
	abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	res := 0
	for i, ss := range s {
		index := charIndexMap[ss]
		res += abs(i - index)
	}
	return res
}

func Test3146(t *testing.T) {
	s := "abc"
	r := "bac"
	res := findPermutationDifference(s, r)
	fmt.Println(res)
}
