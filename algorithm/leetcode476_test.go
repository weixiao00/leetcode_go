package algorithm

import (
	"fmt"
	"testing"
)

func findComplement(num int) int {

	res := 0
	idx := 0
	for num > 0 {
		bin := num & 1
		num >>= 1
		if bin == 0 {
			res = 1<<idx | res
		}
		idx++
	}

	return res
}

func Test476(t *testing.T) {
	num := 5
	res := findComplement(num)
	fmt.Println(res)
}
