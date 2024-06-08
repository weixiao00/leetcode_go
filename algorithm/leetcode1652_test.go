package algorithm

import (
	"fmt"
	"testing"
)

func decrypt(code []int, k int) []int {
	if len(code) == 0 {
		return code
	}

	var afterSum func(k int) int
	afterSum = func(i int) int {
		j := 0
		res := 0
		for j < k {
			i += 1
			if i == len(code) {
				i = 0
			}
			res += code[i]
			j++
		}

		return res
	}

	var beforeSum func(k int) int
	beforeSum = func(i int) int {
		j := 0
		res := 0
		for j > k {
			i -= 1
			if i == -1 {
				i = len(code) - 1
			}
			res += code[i]
			j--
		}

		return res
	}

	res := make([]int, 0, len(code))
	for i := 0; i < len(code); i++ {
		if k == 0 {
			res = append(res, 0)
		} else if k > 0 {
			res = append(res, afterSum(i))
		} else {
			res = append(res, beforeSum(i))
		}
	}

	return res
}

func Test1652(t *testing.T) {
	code := []int{2, 4, 9, 3}
	k := -2
	res := decrypt(code, k)
	fmt.Println(res)
}
