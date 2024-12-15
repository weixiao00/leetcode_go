package algorithm

import (
	"fmt"
	"strconv"
	"testing"
)

func minChanges(n int, k int) int {

	var reverse func(s string) string
	reverse = func(s string) string {
		runes := make([]rune, 0)
		for i := len(s) - 1; i >= 0; i-- {
			runes = append(runes, rune(s[i]))
		}

		return string(runes)
	}

	nStr := strconv.FormatInt(int64(n), 2)
	kStr := strconv.FormatInt(int64(k), 2)
	nStr = reverse(nStr)
	kStr = reverse(kStr)

	length := len(kStr)
	if len(nStr) > len(kStr) {
		length = len(nStr)
	}

	res := 0
	index := 0
	for index < length {
		nIndex := "0"
		if index < len(nStr) {
			nIndex = string(nStr[index])
		}
		kIndex := "0"
		if index < len(kStr) {
			kIndex = string(kStr[index])
		}
		if nIndex == "1" && kIndex == "0" {
			res++
		}
		if nIndex == "0" && kIndex == "1" {
			return -1
		}
		index++
	}

	return res
}

func Test3226(t *testing.T) {
	n := 44
	//fmt.Println(strconv.FormatInt(int64(n), 2))
	k := 2
	res := minChanges(n, k)
	fmt.Println(res)
}
