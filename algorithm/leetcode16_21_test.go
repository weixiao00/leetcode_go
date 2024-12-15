package algorithm

import (
	"fmt"
	"testing"
)

func findSwapValues(array1 []int, array2 []int) []int {

	count1 := 0
	for _, a := range array1 {
		count1 += a
	}
	count2 := 0
	for _, a := range array2 {
		count2 += a
	}

	mp := make(map[int]bool)
	for _, x := range array1 {
		mp[x] = true
	}

	// 如果差值为奇数，不可能出现符合情况的。因为两个数组里都是整数
	// x - y = (count1-count2)/2
	if (count1-count2)&1 == 1 {
		return []int{}
	}
	count := (count1 - count2) / 2
	for _, y := range array2 {
		if mp[count+y] {
			return []int{count + y, y}
		}
	}

	return []int{}
}

func Test16_21(t *testing.T) {
	array1 := []int{1, 2, 3}
	array2 := []int{4, 5, 6}
	res := findSwapValues(array1, array2)
	fmt.Println(res)
}
