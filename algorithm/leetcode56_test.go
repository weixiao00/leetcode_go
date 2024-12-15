package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func merge(intervals [][]int) [][]int {

	if len(intervals) == 0 {
		return [][]int{}
	}

	res := make([][]int, 0)

	// 根据左侧值排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	curMerge := intervals[0]
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval[0] > curMerge[1] {
			res = append(res, curMerge)
			curMerge = interval
			continue
		}
		right := getMax20(interval[1], curMerge[1])
		curMerge = []int{curMerge[0], right}
	}
	res = append(res, curMerge)

	return res
}

func getMax20(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test56(t *testing.T) {
	intervals := [][]int{{1, 4}, {0, 4}}
	res := merge(intervals)
	fmt.Println(res)
}
