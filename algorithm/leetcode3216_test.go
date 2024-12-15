package algorithm

import (
	"fmt"
	"testing"
)

func getSmallestString1(s string) string {

	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		first := runes[i]
		second := runes[i+1]
		if first%2 == 0 && second%2 == 0 && first > second {
			tmp := runes[i]
			runes[i] = runes[i+1]
			runes[i+1] = tmp
			return string(runes)
		} else if first%2 == 1 && second%2 == 1 && first > second {
			tmp := runes[i]
			runes[i] = runes[i+1]
			runes[i+1] = tmp
			return string(runes)
		}
	}

	return s
}

func Test3216(t *testing.T) {
	s := "13"
	res := getSmallestString1(s)
	fmt.Println(res)
}
