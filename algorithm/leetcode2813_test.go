package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func findMaximumElegance(items [][]int, k int) int64 {
	if len(items) == 0 {
		return 0
	}
	res := 0
	// 通过利润进行降序排
	// i, j是索引
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] > items[j][0]
	})

	typeMap := make(map[int]bool, 0)
	profit := 0
	profitReptArr := make([]int, 0)
	for i, item := range items {
		if i < k {
			profit += item[0]
			if !typeMap[item[1]] {
				typeMap[item[1]] = true
			} else {
				profitReptArr = append(profitReptArr, item[0])
			}
		} else {
			if len(profitReptArr) != 0 && !typeMap[item[1]] {
				typeMap[item[1]] = true
				//profitReptArr[len(profitReptArr)-1]最后一个是最小的
				profit += item[0] - profitReptArr[len(profitReptArr)-1]
				profitReptArr = profitReptArr[0 : len(profitReptArr)-1]
			}
		}

		res = getMax10(res, profit+len(typeMap)*len(typeMap))
	}

	return int64(res)
}

func getMax10(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test2813(t *testing.T) {
	//{{8, 4}, {3, 4}, {2, 2}, {1, 3}}
	items := [][]int{{3, 4}, {8, 4}, {2, 2}, {1, 3}}
	k := 2
	elegance := findMaximumElegance(items, k)
	fmt.Println(elegance)
}
