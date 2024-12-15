package algorithm

import (
	"fmt"
	"testing"
)

// 单调栈+贪心
func removeDuplicateLetters(s string) string {

	if len(s) <= 1 {
		return s
	}

	// 可以使用最后元素的索引
	lastIndex := [26]int{}
	for i := range s {
		lastIndex[s[i]-'a'] = i
	}
	visited := [26]bool{}

	var stack []byte
	for i := range s {
		ch := s[i]
		index := ch - 'a'
		if !visited[index] {
			for len(stack) > 0 && stack[len(stack)-1] > ch {
				if lastIndex[stack[len(stack)-1]-'a'] > i {
					visited[stack[len(stack)-1]-'a'] = false
					stack = stack[0 : len(stack)-1]
				} else {
					break
				}
			}
			visited[index] = true
			stack = append(stack, ch)
		}
	}

	return string(stack)
}

func Test316(t *testing.T) {
	//bcabc
	s := "bcabc"
	res := removeDuplicateLetters(s)
	fmt.Println(res)
}
