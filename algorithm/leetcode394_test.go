package algorithm

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

// 栈
// 自己的解法，不优雅
func decodeString(s string) string {

	//"3[a2[c]]"
	leftIndexStack := make([]int, 0)
	numStack := make([]string, 0)
	index := 0
	num := ""
	for true {
		if index >= len(s) {
			return s
		}
		if s[index] >= '0' && s[index] <= '9' {
			num += string(s[index])
		}
		if s[index] == '[' {
			leftIndexStack = append(leftIndexStack, index)
			numStack = append(numStack, num)
			num = ""
		}
		if s[index] == ']' && len(leftIndexStack) > 0 {
			leftIndex := leftIndexStack[len(leftIndexStack)-1]
			leftIndexStack = leftIndexStack[0 : len(leftIndexStack)-1]
			ss := s[leftIndex+1 : index]
			num := numStack[len(numStack)-1]
			atoi, _ := strconv.Atoi(num)
			numStack = numStack[0 : len(numStack)-1]
			s = s[0:leftIndex-len(num)] + strings.Repeat(ss, atoi) + s[index+1:]
			index = leftIndex
			continue
		}
		index++
	}

	return s
}

// 官方题解
func decodeString1(s string) string {

	stack := make([]string, 0)
	ptr := 0
	runes := []rune(s)

	for ptr < len(runes) {
		// 如果是数字
		if unicode.IsDigit(runes[ptr]) {
			builder := strings.Builder{}
			for runes[ptr] != '[' {
				builder.WriteRune(runes[ptr])
				ptr++
			}
			stack = append(stack, builder.String())
			// 如果是字母或者'['
		} else if unicode.IsLetter(runes[ptr]) || runes[ptr] == '[' {
			stack = append(stack, string(runes[ptr]))
			ptr++
			// 是']'
		} else {
			ptr++
			// 不是整个字符串倒续。是所有字符串倒续
			ss := ""
			//builder := strings.Builder{}
			// 字母出来
			for stack[len(stack)-1] != "[" {
				ss = stack[len(stack)-1] + ss
				stack = stack[0 : len(stack)-1]
			}

			// '['出来
			stack = stack[0 : len(stack)-1]

			// 数字出来
			numStr := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			num, _ := strconv.Atoi(numStr)

			// 构造
			repeat := strings.Repeat(ss, num)

			// 加入栈
			stack = append(stack, repeat)
		}
	}

	var builder strings.Builder
	for i := range stack {
		builder.WriteString(stack[i])
	}

	return builder.String()
}

func Test394(t *testing.T) {
	//"zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"
	//"zzzyypqfejkjkfejkjkfejkjkfejkjkyypqfejkjkfejkjkfejkjkfejkjkef"
	s := "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
	res := decodeString1(s)
	fmt.Println(res)
}
