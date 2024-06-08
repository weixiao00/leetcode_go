package algorithm

import (
	"fmt"
	"testing"
)

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges)
	res := make([]int, n+1)
	if n == 0 {
		return res
	}

	// 存储每个节点的入节点信息[1,5] 1是入，5是入度
	inputArr := make([][][]int, n+1)
	for _, edg := range edges {
		input := edg[0]
		output := edg[1]
		weight := edg[2]
		inputArr[input] = append(inputArr[input], []int{output, weight})
		inputArr[output] = append(inputArr[output], []int{input, weight})
	}

	var dfs func(position, root, curr int) int
	dfs = func(position, root, curr int) int {
		res := 0
		if curr == 0 {
			res++
		}
		nodeInputArr := inputArr[position]
		for _, arr := range nodeInputArr {
			if root != arr[0] {
				res += dfs(arr[0], position, (arr[1]+curr)%signalSpeed)
			}
		}

		return res
	}

	for i := 0; i < n+1; i++ {
		pre := 0
		nodeInputArr := inputArr[i]
		// 计算左右可以联通的节点数量乘积
		for _, arr := range nodeInputArr {
			last := dfs(arr[0], i, arr[1]%signalSpeed)
			res[i] += pre * last
			pre += last
		}
	}

	return res
}

func dfs2(position, root, signalSpeed, curr int, inputArr [][][]int) int {
	res := 0
	if curr == 0 {
		res++
	}
	nodeInputArr := inputArr[position]
	for _, arr := range nodeInputArr {
		if root != arr[0] {
			res += dfs2(arr[0], position, signalSpeed, arr[1]%signalSpeed, inputArr)
		}
	}

	return res
}

func Test3067(t *testing.T) {
	edges := [][]int{{0, 6, 3}, {6, 5, 3}, {0, 3, 1}, {3, 2, 7}, {3, 1, 6}, {3, 4, 2}}
	signalSpeed := 3
	servers := countPairsOfConnectableServers(edges, signalSpeed)
	//[2,0,0,0,0,0,2]
	fmt.Println(servers)
}
