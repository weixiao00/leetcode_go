package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func maxIncreasingCells(mat [][]int) int {

	if len(mat) == 0 {
		return 0
	}

	m := len(mat)
	n := len(mat[0])
	row := make([]int, m)
	column := make([]int, n)
	// 存储每个元素的值和索引
	mp := make(map[int][][]int, 0)

	for i := range mat {
		for j := range mat[i] {
			val, ok := mp[mat[i][j]]
			if !ok {
				val = make([][]int, 0)
			}
			val = append(val, []int{i, j})
			// 这行必须得加，因为不是指针
			mp[mat[i][j]] = val
		}
	}

	// 闭包
	//var sortMap func(mp map[int][][]int) map[int][][]int
	//sortMap = func(mp map[int][][]int) map[int][][]int {
	//	keys := make([]int, 0)
	//	for key := range mp {
	//		keys = append(keys, key)
	//	}
	//	sort.Ints(keys)
	//	newMp := make(map[int][][]int, 0)
	//	for _, key := range keys {
	//		newMp[key] = mp[key]
	//	}
	//	return newMp
	//}

	// 通过key升序排
	//mp = sortMap(mp)

	keys := make([]int, 0)
	for key := range mp {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// 不能直接循环map。每个打印的顺序都是不一样的
	for _, key := range keys {
		val := mp[key]
		compRes := make([]int, 0)
		// 对于相同的元素，不能相互影响
		for _, arr := range val {
			rowIdx := arr[0]
			columnIdx := arr[1]
			// 行最大或者列最大加1
			compRes = append(compRes, getMax17(row[rowIdx], column[columnIdx])+1)
		}

		// 更新行最大值和列最大值
		for i, arr := range val {
			rowIdx := arr[0]
			columnIdx := arr[1]
			compVal := compRes[i]
			// 这里取最大值，是有行数据相等的情况
			row[rowIdx] = getMax17(row[rowIdx], compVal)
			column[columnIdx] = getMax17(column[columnIdx], compVal)
		}
	}

	res := 0
	for _, max := range row {
		res = getMax17(res, max)
	}

	return res
}

func getMax17(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test2713(t *testing.T) {
	mat := [][]int{{1, -8}, {4, 4}, {-3, -9}}
	res := maxIncreasingCells(mat)
	fmt.Println(res)
}

func Test3333(t *testing.T) {
	//mutex := sync.Mutex{}
	//mutex.Lock()
	mp := make(map[int]int, 0)
	mp[1] = 1
	mp[2] = 2
	mp[3] = 3
	mp[4] = 4
	mp[5] = 5
	//mutex.Unlock()
	//直接遍历
	for key, val := range mp {
		fmt.Println(key, val)
	}
}
