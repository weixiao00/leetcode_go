package algorithm

import (
	"fmt"
	"testing"
)

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {

	n := len(values)
	// 还未初始化
	graphs := make([][][]int, n)

	// 每个点入度的初始化
	for _, edge := range edges {
		inputArr0 := graphs[edge[0]]
		if inputArr0 == nil {
			inputArr0 = make([][]int, 0)
		}
		inputArr0 = append(inputArr0, []int{edge[1], edge[2]})
		graphs[edge[0]] = inputArr0
		inputArr1 := graphs[edge[1]]
		if inputArr1 == nil {
			inputArr1 = make([][]int, 0)
		}
		inputArr1 = append(inputArr1, []int{edge[0], edge[2]})
		graphs[edge[1]] = inputArr1
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := values[0]
	var dfs func(curTime, curValue, prevIdx int, visited []bool)
	dfs = func(curTime, curValue, prevIdx int, visited []bool) {
		// 递归结束条件
		if curTime > maxTime {
			return
		}

		if prevIdx == 0 {
			res = getMax(res, curValue)
		}

		for _, graph := range graphs[prevIdx] {
			if visited[graph[0]] {
				dfs(curTime+graph[1], curValue, graph[0], visited)
			} else {
				visited[graph[0]] = true
				dfs(curTime+graph[1], curValue+values[graph[0]], graph[0], visited)
				visited[graph[0]] = false
			}
		}
	}

	visited := make([]bool, n)
	visited[0] = true
	dfs(0, values[0], 0, visited)
	return res
}

func Test2065(t *testing.T) {
	values := []int{0, 32, 10, 43}
	edges := [][]int{{0, 1, 10}, {1, 2, 15}, {0, 3, 10}}
	maxTime := 49
	quality := maximalPathQuality(values, edges, maxTime)
	fmt.Println(quality)
}
