package algorithm

import (
	"fmt"
	"testing"
)

// 没过
func numberOfSets(n int, maxDistance int, roads [][]int) int {

	if n == 0 {
		return 0
	}

	graph := make([][][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([][]int, 0)
	}
	for _, road := range roads {
		graph[road[0]] = append(graph[road[0]], []int{road[1], road[2]})
		graph[road[1]] = append(graph[road[1]], []int{road[0], road[2]})
	}

	res := 0
	// 表示已经访问过
	visited := make(map[int]bool)
	var dfs func(curDistance, idx int)
	dfs = func(curDistance, idx int) {

		if curDistance > maxDistance {
			return
		}

		visited[idx] = true
		if len(visited) == n {
			res++
		}

		for _, node := range graph[idx] {
			if visited[node[0]] {
				continue
			}
			// 留
			dfs(curDistance+node[1], node[0])
			// 不留
			dfs(curDistance, node[0])
		}

		//delete(visited, idx)
	}

	for _, node := range graph[0] {
		//visited[node[0]] = true
		// 留
		dfs(node[1], 0)
		// 不留
		//delete(visited, node[0])
	}
	dfs(0, 0)

	return res
}

func Test1959(t *testing.T) {
	n := 3
	maxDistance := 5
	roads := [][]int{{0, 1, 2}, {1, 2, 10}, {0, 2, 10}}
	res := numberOfSets(n, maxDistance, roads)
	fmt.Println(res)
}
