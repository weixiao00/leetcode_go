package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

// 就是排序的时间复杂度
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	merge := make([][]int, 0, len(difficulty))
	for i, _ := range difficulty {
		merge = append(merge, []int{difficulty[i], profit[i]})
	}
	sort.Slice(merge, func(i, j int) bool {
		return merge[i][0] < merge[j][0]
	})
	sort.Ints(worker)

	len := len(merge)
	res := 0
	best := 0
	j := 0
	for _, work := range worker {
		for j < len && work >= merge[j][0] {
			best = getMax3(merge[j][1], best)
			j++
		}
		res += best
	}

	return res
}

func getMax3(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Test826(t *testing.T) {
	difficulty := []int{2, 4, 6, 8, 10}
	profit := []int{10, 20, 30, 40, 50}
	worker := []int{4, 5, 6, 7}
	assignment := maxProfitAssignment(difficulty, profit, worker)
	fmt.Println(assignment)
}
