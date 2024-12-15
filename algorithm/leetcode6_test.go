package algorithm

import (
	"fmt"
	"strings"
	"testing"
)

// 自己的解法，找规律，但是不优雅
func convert(s string, numRows int) string {

	if len(s) == 0 {
		return s
	}

	if numRows == 1 {
		return s
	}

	if len(s) <= numRows {
		return s
	}

	res := ""
	for i := 0; i < numRows; i++ {

		j := i
		// 第一行和最后一行
		if i == 0 || i == numRows-1 {
			for j < len(s) {
				res += string(s[j])
				j += numRows + numRows - 2
			}
			continue
		}

		// 中间行
		res += string(s[j])
		k := 0
		for true {
			k += numRows + numRows - 2
			if k-i >= len(s) {
				break
			}
			res += string(s[k-i])
			if k+i < len(s) {
				res += string(s[k+i])
			}
		}
	}

	return res
}

// 官方题解，flag变换
func convert1(s string, numRows int) string {

	if len(s) == 0 {
		return s
	}
	if numRows < 2 {
		return s
	}

	strs := make([]string, numRows)

	i := 0
	flag := -1
	for _, ss := range s {
		strs[i] += string(ss)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}

	return strings.Join(strs, "")
}

func Test6(t *testing.T) {

	s := "PAYPALISHIRING"
	numRows := 3
	res := convert1(s, numRows)
	fmt.Println(res)
	//PAHNAPLSIIGYIR
}
