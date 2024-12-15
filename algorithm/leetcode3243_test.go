package algorithm

import (
	"container/list"
	"fmt"
	"testing"
)

// bfs.类似Dijkstra算法
func shortestDistanceAfterQueries(n int, queries [][]int) []int {

	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < n-1; i++ {
		graph[i] = append(graph[i], i+1)
	}

	var bfs func() int
	bfs = func() int {
		dict := make([]int, n)

		queue := list.New()
		queue.PushBack(0)
		for queue.Len() > 0 {
			front := queue.Front()
			queue.Remove(front)
			x := front.Value.(int)
			for _, y := range graph[x] {
				if dict[y] > 0 {
					continue
				}
				queue.PushBack(y)
				dict[y] = dict[x] + 1
			}
		}

		return dict[n-1]
	}

	answer := make([]int, len(queries))
	for i := range queries {
		graph[queries[i][0]] = append(graph[queries[i][0]], queries[i][1])
		answer[i] = bfs()
	}

	return answer
}

func Test3243(t *testing.T) {
	n := 5
	queries := [][]int{{2, 4}, {0, 2}, {0, 4}}
	answer := shortestDistanceAfterQueries(n, queries)
	fmt.Println(answer)
}
