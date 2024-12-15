package algorithm

import (
	"fmt"
	"testing"
)

// 动态规划
func maxRepeating(sequence string, word string) int {

	if len(word) > len(sequence) {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([]int, len(sequence)+1)

	res := 0
	for i := 0; i < len(sequence)-len(word)+1; i++ {
		if word == sequence[i:i+len(word)] {
			dp[i+len(word)] = dp[i] + 1
			res = getMax(res, dp[i+len(word)])
		}
	}

	return res
}

// 直接暴力
func maxRepeating1(sequence string, word string) int {
	if len(word) > len(sequence) {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	index := 0
	for i := range sequence {
		curMax := 0
		j := i
		for j < len(sequence) && sequence[j] == word[index] {
			j++
			index++

			if index >= len(word) {
				index = 0
				curMax++
			}
		}
		index = 0
		res = getMax(res, curMax)
	}

	return res
}

func Test1668(t *testing.T) {
	sequence := "aaa"
	word := "aa"
	res := maxRepeating1(sequence, word)
	fmt.Println(res)
}
