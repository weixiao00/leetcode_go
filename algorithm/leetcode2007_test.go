package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func findOriginalArray(changed []int) []int {
	if len(changed) == 0 {
		return nil
	}

	sort.Ints(changed)
	countMap := make(map[int]int, len(changed))

	for _, c := range changed {
		countMap[c] = countMap[c] + 1
	}

	res := make([]int, 0)

	for _, c := range changed {
		if countMap[c] == 0 {
			continue
		}
		countMap[c] = countMap[c] - 1
		val := countMap[c*2]
		if val == 0 {
			return []int{}
		}
		res = append(res, c)
		countMap[c*2] = val - 1
	}

	return res
}

func Test2007(t *testing.T) {
	arr := []int{0}
	champion := findOriginalArray(arr)
	fmt.Println(champion)
}
