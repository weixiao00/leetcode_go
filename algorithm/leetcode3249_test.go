package algorithm

import (
	"fmt"
	"testing"
)

// dfs。得传入父节点
func countGoodNodes(edges [][]int) int {

	if len(edges) == 0 {
		return 0
	}

	ingree := make([][]int, len(edges)+1)
	for i := range ingree {
		ingree[i] = make([]int, 0)
	}

	for _, edge := range edges {
		v0 := edge[0]
		v1 := edge[1]
		ingree[v0] = append(ingree[v0], v1)
		ingree[v1] = append(ingree[v1], v0)
	}

	res := 0
	var dfs func(node, parent int) int
	dfs = func(node, parent int) int {
		count := 0
		valid := true
		total := 1
		for _, out := range ingree[node] {
			if out == parent {
				continue
			}
			subCount := dfs(out, node)
			if count == 0 {
				count = subCount
			} else if count != subCount {
				valid = false
			}
			total += subCount
		}
		if valid {
			res++
		}
		return total
	}

	dfs(0, -1)

	return res
}

func Test3249(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {0, 5}, {1, 6}, {2, 7}, {3, 8}}
	res := countGoodNodes(edges)
	fmt.Println(res)
}
