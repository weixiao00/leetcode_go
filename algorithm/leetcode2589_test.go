package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

// 尽可能利用之前的时间
func findMinimumTime(tasks [][]int) int {

	if len(tasks) == 0 {
		return 0
	}
	res := 0
	n := len(tasks)
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})
	run := make([]int, tasks[n-1][1]+1)
	for _, task := range tasks {
		start := task[0]
		end := task[1]
		duration := task[2]
		for i := start; i <= end; i++ {
			if run[i] == 1 {
				duration--
			}
		}
		//res += getMax1(0, duration)
		// 还有一些之前没有运行的
		// 从后往前，是因为。tasks是按照task[i][1]递增的。尽可能在下一个里边被使用到
		for i := end; i >= start && duration > 0; i-- {
			if run[i] == 0 {
				duration--
				run[i] = 1
			}
		}
	}
	for _, r := range run {
		res += r
	}
	return res
}

func getMax1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test2589(t *testing.T) {
	tasks := [][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}}
	time := findMinimumTime(tasks)
	fmt.Println(time)
}
