package algorithm

import (
	"fmt"
	"strconv"
	"testing"
)

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	xstr := strconv.Itoa(x)

	left := 0
	right := len(xstr) - 1
	for left < right {
		if xstr[left] != xstr[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func Test9(t *testing.T) {
	x := -1210
	res := isPalindrome(x)
	fmt.Println(res)
}
