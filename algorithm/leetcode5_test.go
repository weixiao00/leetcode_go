package algorithm

import (
	"fmt"
	"testing"
)

// 最长回文字符串。单双串处理
func longestPalindrome(s string) string {

	if len(s) == 0 {
		return s
	}

	var extend func(s string, left, right int) []int
	extend = func(s string, left, right int) []int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return []int{left + 1, right - 1}
	}

	res := make([]int, 2)
	for i := range s {
		arr1 := extend(s, i, i)
		arr2 := extend(s, i, i+1)
		if arr1[1]-arr1[0] > arr2[1]-arr2[0] {
			if arr1[1]-arr1[0] > res[1]-res[0] {
				res[1] = arr1[1]
				res[0] = arr1[0]
			}
		} else {
			if arr2[1]-arr2[0] > res[1]-res[0] {
				res[1] = arr2[1]
				res[0] = arr2[0]
			}
		}
	}

	return s[res[0] : res[1]+1]
}

// 动态规划解法
func longestPalindrome1(s string) string {

	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, len(s))

	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	// 初始化dp，一个元素肯定回文
	for i := range s {
		dp[i][i] = true
	}

	res := 1
	begin := 0
	// 长度为2的判断是否回文
	for length := 2; length <= len(s); length++ {
		for i := 0; i < len(s); i++ {
			// 右坐标
			j := length + i - 1
			if j >= len(s) {
				break
			}

			if s[i] != s[j] {
				continue
			}

			if length <= 3 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}

			if dp[i][j] && res < length {
				res = length
				begin = i
			}
		}
	}

	return s[begin : begin+res]
}

func getMax22(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test5(t *testing.T) {
	s := "bb"
	res := longestPalindrome1(s)
	fmt.Println(res)
}
