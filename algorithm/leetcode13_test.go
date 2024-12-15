package algorithm

import (
	"fmt"
	"testing"
)

// 自己的解法
func romanToInt(s string) int {

	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	vals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := 0
	if len(s) == 0 {
		return 0
	}
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		val := vals[i]

		if len(s) == 0 {
			return res
		}

		for len(s) > 0 {
			first := ""
			if len(s) > 0 {
				first = s[0:1]
			}
			if first == val {
				res += key
				s = s[1:]
			} else {
				break
			}
		}

		for len(s) > 1 {
			second := ""
			if len(s) > 1 {
				second = s[0:2]
			}
			if second == val {
				res += key
				s = s[2:]
			} else {
				break
			}
		}
	}

	return res
}

// 官方解法
func romanToInt1(s string) int {

	mp := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	res := 0
	for i := 1; i < len(s); i++ {
		if mp[rune(s[i-1])] < mp[rune(s[i])] {
			res -= mp[rune(s[i-1])]
		} else {
			res += mp[rune(s[i-1])]
		}
	}

	// 最后一个数一定是加
	res += mp[rune(s[len(s)-1])]
	return res
}

func Test13(t *testing.T) {
	s := "MCMXCIV"
	res := romanToInt1(s)
	fmt.Println(res)
}
