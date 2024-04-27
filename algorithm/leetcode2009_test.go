package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func Test2009(t *testing.T) {
	arr := []int{8, 5, 9, 9, 8, 4}
	fmt.Println(minOperations(arr))
}

func minOperations(nums []int) int {
	if nums == nil {
		return 0
	}
	numMap := make(map[int]struct{}, 0)
	for _, n := range nums {
		numMap[n] = struct{}{}
	}
	length := len(nums)
	res := length

	values := make([]int, 0, len(numMap))
	for key, _ := range numMap {
		values = append(values, key)
	}

	sort.Ints(values)
	j := 0
	for i := 0; i < len(values); i++ {
		min := values[i]
		max := min + length - 1
		for j < len(values) && values[j] <= max {
			res = getMin(length-(j-i+1), res)
			j++
		}
	}

	return res
}

func getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
