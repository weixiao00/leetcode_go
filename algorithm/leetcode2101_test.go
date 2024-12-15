package algorithm

import (
	"fmt"
	"testing"
)

// 并查集解决的无向图
// 并查集试试？？？
// dfs是递归实现，bfs是队列实现

// 过了, dfs
func maximumDetonation(bombs [][]int) int {
	if len(bombs) == 0 {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var sqrt func(a int) int
	sqrt = func(a int) int {
		return a * a
	}

	var canBom func(preBom, curBom []int) bool
	canBom = func(preBom, curBom []int) bool {
		return sqrt(preBom[0]-curBom[0])+sqrt(preBom[1]-curBom[1]) <= sqrt(preBom[2])
	}

	visited := make([]bool, len(bombs))
	res := 1
	var dfs func(preIdx, preCount int) int
	dfs = func(preIdx, preCount int) int {
		preBom := bombs[preIdx]
		for i, bom := range bombs {
			if visited[i] {
				continue
			}
			// 坐标半径
			if !canBom(preBom, bom) {
				continue
			}
			visited[i] = true
			preCount = dfs(i, preCount+1)
		}

		return preCount
	}

	for i := range bombs {
		// 重新设置是否访问
		visited = make([]bool, len(bombs))
		visited[i] = true
		res = getMax(res, dfs(i, 1))
	}

	return res
}

// 官方解法 bfs
// 图论的解法。可以用dfs和bfs遍历。图的最短路径使用dijkstra算法
func maximumDetonation1(bombs [][]int) int {
	if len(bombs) == 0 {
		return 0
	}
	var sqrt func(a int) int
	sqrt = func(a int) int {
		return a * a
	}

	var canBom func(preBom, curBom []int) bool
	canBom = func(preBom, curBom []int) bool {
		return sqrt(preBom[0]-curBom[0])+sqrt(preBom[1]-curBom[1]) <= sqrt(preBom[2])
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 维护引爆关系
	mp := make(map[int][]int)
	for i := range bombs {
		for j := range bombs {
			if i != j && canBom(bombs[i], bombs[j]) {
				val, ok := mp[i]
				if !ok {
					val = make([]int, 0)
				}
				val = append(val, j)
				mp[i] = val
			}
		}
	}

	res := 1
	for i := range bombs {
		// 剪枝
		visited := make([]bool, len(bombs))
		queue := make([]int, 0)
		queue = append(queue, i)
		curRes := 1
		visited[i] = true
		for len(queue) != 0 {
			first := queue[0]
			queue = queue[1:]

			for _, j := range mp[first] {
				if visited[j] {
					continue
				}
				queue = append(queue, j)
				curRes++
				visited[j] = true
			}
		}

		res = getMax(res, curRes)
	}

	return res
}

func Test2101(t *testing.T) {
	bombs := [][]int{{855, 82, 158}, {17, 719, 430}, {90, 756, 164}, {376, 17, 340}, {691, 636, 152},
		{565, 776, 5}, {464, 154, 271}, {53, 361, 162}, {278, 609, 82}, {202, 927, 219}, {542, 865, 377},
		{330, 402, 270}, {720, 199, 10}, {986, 697, 443}, {471, 296, 69}, {393, 81, 404}, {127, 405, 177}}
	//bombs := [][]int{{2, 1, 3}, {6, 1, 4}}
	res := maximumDetonation(bombs)
	fmt.Println(res)
}
