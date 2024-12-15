package algorithm

import (
	"fmt"
	"testing"
)

//dfs
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	mp := make(map[rune][]string)
	mp['2'] = []string{"a", "b", "c"}
	mp['3'] = []string{"d", "e", "f"}
	mp['4'] = []string{"g", "h", "i"}
	mp['5'] = []string{"j", "k", "l"}
	mp['6'] = []string{"m", "n", "o"}
	mp['7'] = []string{"p", "q", "r", "s"}
	mp['8'] = []string{"t", "u", "v"}
	mp['9'] = []string{"w", "x", "y", "z"}

	res := make([]string, 0)
	var dfs func(idx int, mergeStr string)
	dfs = func(idx int, mergeStr string) {
		if idx == len(digits) {
			res = append(res, mergeStr)
			return
		}
		u := digits[idx]
		chars := mp[rune(u)]
		for _, str := range chars {
			dfs(idx+1, mergeStr+str)
		}
	}

	dfs(0, "")

	return res
}

func Test17(t *testing.T) {
	digits := "234"
	res := letterCombinations(digits)
	fmt.Println(res)
}
