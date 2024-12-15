package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

// hash函数
func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {

	if len(moveFrom) == 0 {
		return nums
	}

	idxCountMap := make(map[int]bool)

	for _, num := range nums {
		idxCountMap[num] = true
	}

	for i := range moveFrom {
		from := moveFrom[i]
		to := moveTo[i]
		idxCountMap[from] = false
		idxCountMap[to] = true
	}

	res := make([]int, 0)
	for idx, count := range idxCountMap {
		if count == false {
			continue
		}
		res = append(res, idx)
	}
	sort.Ints(res)
	return res
}

func Test2766(t *testing.T) {
	nums := []int{1, 1, 3, 3}
	moveFrom := []int{1, 3}
	moveTo := []int{2, 2}
	res := relocateMarbles(nums, moveFrom, moveTo)
	fmt.Println(res)
	//2
}
