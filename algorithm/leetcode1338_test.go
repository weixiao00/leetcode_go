package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

// hash存储数量
// 贪心
func minSetSize(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	mp := make(map[int]int)
	for i := range arr {
		mp[arr[i]]++
	}

	keyValCount := make([][]int, 0)
	for key, val := range mp {
		keyValCount = append(keyValCount, []int{key, val})
	}
	sort.Slice(keyValCount, func(i, j int) bool {
		return keyValCount[i][1] > keyValCount[j][1]
	})

	res := 0
	count := 0
	for i := range keyValCount {
		count += keyValCount[i][1]
		res++
		if count >= len(arr)/2 {
			break
		}
	}

	return res
}

func Test1338(t *testing.T) {
	arr := []int{7, 7, 7, 7, 7, 7}
	res := minSetSize(arr)
	fmt.Println(res)
}
