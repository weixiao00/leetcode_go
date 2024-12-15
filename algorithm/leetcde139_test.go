package algorithm

import (
	"fmt"
	"strings"
	"testing"
)

// 动态规划+记忆搜索(一纬动态规划)
func wordBreak1(s string, wordDict []string) bool {

	if len(s) == 0 {
		return true
	}

	var contain func(wordDict []string, subStr string) bool
	contain = func(wordDict []string, subStr string) bool {
		for _, word := range wordDict {
			if word == subStr {
				return true
			}
		}
		return false
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	// 求每个dp[i]
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if contain(wordDict, s[j:i]) && dp[j] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}

// dfs超时了
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	var dfs func(curStr string) bool
	dfs = func(curStr string) bool {
		if len(curStr) == 0 {
			return true
		}
		res := false
		for idx := 0; idx < len(wordDict); idx++ {
			word := wordDict[idx]
			if strings.HasPrefix(curStr, word) {
				subStr := strings.Replace(curStr, word, "", 1)
				res = res || dfs(subStr)
			}
		}
		return res
	}

	return dfs(s)
}

func Test139(t *testing.T) {
	s := "catskicatcats"
	wordDict := []string{"cats", "cat", "dog", "ski"}
	res := wordBreak1(s, wordDict)
	fmt.Println(res)
}
