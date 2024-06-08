package algorithm

import (
	"fmt"
	"testing"
)

func missingRolls(rolls []int, mean int, n int) []int {
	rollCount := 0

	for _, roll := range rolls {
		rollCount += roll
	}
	len := len(rolls)

	totalCount := (n + len) * mean

	nCount := totalCount - rollCount

	res := make([]int, 0)
	if n*6 < nCount {
		return res
	}
	if n*1 > nCount {
		return res
	}
	div := nCount / n
	rem := nCount % n

	idx := 0
	for idx < n {
		if rem > 0 {
			rem--
			res = append(res, div+1)
			idx++
			continue
		}
		res = append(res, div)
		idx++
	}

	return res
}

func Test2028(t *testing.T) {
	rolls := []int{1, 2, 3, 4}
	mean := 6
	n := 4
	arr := missingRolls(rolls, mean, n)
	fmt.Println(arr)
}
