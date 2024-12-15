package algorithm

import (
	"fmt"
	"testing"
)

// 动态规划+dfs
func partition(s string) [][]string {

	if len(s) == 0 {
		return [][]string{}
	}

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	// 全部填充为true。
	for i := 0; i < len(dp); i++ {
		dp[i][i] = true
	}

	var palindromic func()
	palindromic = func() {
		length := 1
		for length <= len(s) {
			for left := 0; left < len(s); left++ {
				right := left + length
				if right >= len(s) {
					break
				}
				if s[left] != s[right] {
					continue
				}
				if length <= 2 {
					dp[left][right] = true
					continue
				}
				dp[left][right] = dp[left+1][right-1]
			}
			length++
		}
	}

	// i，j判断是否是回文串
	palindromic()

	// 搜索
	res := make([][]string, 0)
	var dfs func(i int, ans []string)
	dfs = func(i int, ans []string) {
		if i == len(s) {
			tmp := make([]string, len(ans))
			copy(tmp, ans)
			res = append(res, tmp)
			return
		}

		for j := i; j < len(s); j++ {
			if dp[i][j] {
				ans = append(ans, s[i:j+1])
				dfs(j+1, ans)
				ans = ans[0 : len(ans)-1]
			}
		}
	}

	dfs(0, []string{})
	return res
}

func Test131(t *testing.T) {
	//[["a","b","b","a","b"],["a","b","bab"],["a","bb","a","b"],["a","bbab"],["abba","b"]]
	s := "abbab"
	res := partition(s)
	fmt.Println(res)
}
