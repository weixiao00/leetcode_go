package algorithm

import (
	"fmt"
	"strings"
	"testing"
)

// 贪心
// 分出来尽可能多的子字符串
func partitionLabels(s string) []int {

	list := make([]string, 0)
	var merge func(index int, char string)
	merge = func(index int, char string) {
		builder := strings.Builder{}
		newList := make([]string, 0)
		for i := range list {
			if i < index {
				newList = append(newList, list[i])
				continue
			}
			builder.WriteString(list[i])
		}
		builder.WriteString(char)
		newList = append(newList, builder.String())
		list = newList
	}

	for i := range s {
		char := string(s[i])
		index := -1
		for j := range list {
			if strings.Contains(list[j], char) {
				index = j
				break
			}
		}
		if index >= 0 {
			merge(index, char)
			continue
		}
		list = append(list, char)
	}

	res := make([]int, 0)
	for i := range list {
		res = append(res, len(list[i]))
	}

	return res
}

// 官方解法
func partitionLabels1(s string) []int {

	last := [26]int{}
	for i := range s {
		last[s[i]-'a'] = i
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := make([]int, 0)
	start := 0
	end := 0
	for i := range s {
		end = getMax(end, last[s[i]-'a'])
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}

	return res
}

func Test763(t *testing.T) {
	s := "ababcbacadefegdehijhklij"
	res := partitionLabels1(s)
	fmt.Println(res)
}
