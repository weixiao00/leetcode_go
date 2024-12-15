package algorithm

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 删除开始的结尾的空格
	s = strings.Trim(s, " ")

	res := 0
	symbol := 1
	for i, sr := range s {
		if i == 0 {
			if (sr < '0' || sr > '9') && (sr != '+' && sr != '-') {
				return 0
			}
			if sr == '-' {
				symbol = -1
				continue
			} else if sr == '+' {
				symbol = 1
				continue
			} else if sr == '0' {
				continue
			}
		}

		if sr < '0' || sr > '9' {
			return res
		}

		if symbol == 1 {
			res = res*10 + int(sr-'0')
		} else {
			res = res*10 - int(sr-'0')
		}
		if res > math.MaxInt32 {
			return math.MaxInt32
		}
		if res < math.MinInt32 {
			return math.MinInt32
		}
	}

	return res
}

func Test8(t *testing.T) {
	s := "-91283472332"
	fmt.Println(myAtoi(s))
}
