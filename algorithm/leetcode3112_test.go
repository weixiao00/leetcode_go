package algorithm

import (
	"fmt"
	"testing"
)

// dfs超时了，牛X
// 需要用到优先队列，使用java实现了
// dijkstra算法
func minimumTime(n int, edges [][]int, disappear []int) []int {

	if n == 0 {
		return []int{}
	}

	graph := make([][][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([][]int, 0)
	}

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], []int{edge[1], edge[2]})
		graph[edge[1]] = append(graph[edge[1]], []int{edge[0], edge[2]})
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	visited := make(map[int]bool)
	var dfs func(curTime, idx int)
	dfs = func(curTime, idx int) {

		if visited[idx] || curTime >= disappear[idx] || (res[idx] != -1 && res[idx] < curTime) {
			return
		}

		res[idx] = curTime
		visited[idx] = true
		for _, node := range graph[idx] {
			// 应该得是先计算最短的距离的节点。要不这样会超时
			dfs(curTime+node[1], node[0])
		}
		visited[idx] = false
	}

	res[0] = 0
	dfs(0, 0)
	return res
}

func Test3112(t *testing.T) {
	//n = 3, edges = [[0,1,2],[1,2,1],[0,2,4]], disappear = [1,1,5]
	n := 3
	edges := [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}}
	disappear := []int{1, 1, 5}
	res := minimumTime(n, edges, disappear)
	fmt.Println(res)
}
