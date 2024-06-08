package algorithm

import (
	"fmt"
	"testing"
)

func minimumSteps(s string) int64 {

	left := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			res += i - left
			left++
		}
	}

	return int64(res)
}

// leetcode官方解法
func minimumSteps1(s string) int64 {

	res := 0
	oneTimes := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			oneTimes++
		} else {
			res += oneTimes
		}
	}

	return int64(res)
}

func Test2938(t *testing.T) {
	s := "110"
	res := minimumSteps1(s)
	fmt.Println(res)
}
