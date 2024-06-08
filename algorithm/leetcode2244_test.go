package algorithm

import (
	"fmt"
	"testing"
)

func minimumRounds(tasks []int) int {

	if len(tasks) == 0 {
		return 0
	}
	eleMap := make(map[int]int, 0)
	for _, t := range tasks {
		eleMap[t] += 1
	}

	res := 0
	for _, val := range eleMap {
		if val == 1 {
			return -1
		}
		div := val / 3
		rem := val % 3
		res += div
		if rem != 0 {
			res += 1
		}
	}

	return res
}

func Test2244(t *testing.T) {
	tasks := []int{2, 3, 3}
	refill := minimumRounds(tasks)
	fmt.Println(refill)
}

func TestMap(t *testing.T) {
	m := make(map[int]int, 0)
	m[1] = 1
	m[3] = 3
	m[2] = 2
	for i := range m {
		fmt.Println(i)
	}
}
