package algorithm

import (
	"fmt"
	"testing"
)

func minDistance(word1 string, word2 string) int {

	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	l1 := len(word1)
	l2 := len(word2)
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// 空字符串和非空，更改的个数就是非空
	for i := 0; i <= l2; i++ {
		dp[0][i] = i
	}
	for j := 0; j <= l1; j++ {
		dp[j][0] = j
	}

	// dp表示前i个字符和前j个字符的编辑距离
	// A插入一个字符
	// B插入一个字符
	// A修改一个字符(B修改一个字符)
	//if (word1.charAt(i-1) == word2.charAt(j-1)) {
	//                    dp[i][j] = dp[i-1][j-1];
	//                } else {
	//                    dp[i][j] = Math.min(Math.min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])+1;
	//                }
	// 过了
	// dp[i-1][j] w1插入一个字符
	// dp[i][j-1] w2插入一个字符
	// dp[i-1][j-1] w1修改一个字符
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] != word2[j-1] {
				dp[i][j] = getMin(getMin(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1]+1)
			} else {
				dp[i][j] = getMin(getMin(dp[i-1][j]+1, dp[i][j-1]+1), dp[i-1][j-1])
			}
		}
	}

	return dp[l1][l2]
}

func Test72(t *testing.T) {
	word1 := "pneumonoultramicroscopicsilicovolcanoconiosis"
	word2 := "ultramicroscopically"
	res := minDistance(word1, word2)
	fmt.Println(res)
}
