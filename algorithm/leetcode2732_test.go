package algorithm

import (
	"fmt"
	"testing"
)

// 数组里只包含0，1
func goodSubsetofBinaryMatrix(grid [][]int) []int {
	if len(grid) == 0 {
		return []int{}
	}

	m := len(grid)
	n := len(grid[0])

	// 返回的数据最多一行或者两行
	res := make([]int, 0)
	mp := make(map[int]int, 0)
	for i := 0; i < m; i++ {
		bir := 0
		for j := 0; j < n; j++ {
			bir |= grid[i][j] << j
		}
		mp[bir] = i
	}

	// 包含0。直接满足
	i, ok := mp[0]
	if ok {
		res = append(res, i)
		return res
	}

	keys := make([]int, 0)
	for key, _ := range mp {
		keys = append(keys, key)
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for _, key1 := range keys {
		for _, key2 := range keys {
			if key1&key2 == 0 {
				res = append(res, getMin(mp[key1], mp[key2]))
				res = append(res, getMax(mp[key1], mp[key2]))
				return res
			}
		}
	}

	return res
}

func Test2732(t *testing.T) {
	grid := [][]int{{0, 1, 1, 0}, {0, 0, 0, 1}, {1, 1, 1, 1}}
	res := goodSubsetofBinaryMatrix(grid)
	fmt.Println(res)
}
