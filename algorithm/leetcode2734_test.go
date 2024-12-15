package algorithm

import (
	"fmt"
	"testing"
)

// 获取最小的字符串。找每个字符的前一个字符
func smallestString(s string) string {
	if len(s) == 0 {
		return s
	}

	bytes := []byte(s)
	var findNotA func(bytes []byte) int
	findNotA = func(bytes []byte) int {

		for i, b := range bytes {
			if b != 'a' {
				return i
			}
		}
		return len(bytes)
	}
	var findAfterA func(bytes []byte, idx int) int
	findAfterA = func(bytes []byte, idx int) int {

		for ; idx < len(bytes); idx++ {
			if bytes[idx] == 'a' {
				return idx
			}
		}
		return len(bytes)
	}

	res := make([]byte, 0)
	notA := findNotA(bytes)
	// 全部是a。还必须得操作一次
	if notA == len(bytes) {
		bytes[len(bytes)-1] = 'z'
		return string(bytes)
	}

	afterA := findAfterA(bytes, notA)
	for i, b := range bytes {
		if i >= notA && i < afterA {
			res = append(res, b-1)
		} else {
			res = append(res, b)
		}
	}

	return string(res)
}

func Test2734(t *testing.T) {
	s := "leetcode"
	res := smallestString(s)
	fmt.Println(res)
}
