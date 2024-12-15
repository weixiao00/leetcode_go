package algorithm

import (
	"fmt"
	"testing"
)

// 模拟
func numberOfAlternatingGroups1(colors []int, k int) int {

	cnt := 1
	res := 0

	n := len(colors)
	// 往前几步。因为是访问索引为0的时候直接判断
	for i := -k + 2; i < len(colors); i++ {
		if colors[(i+n)%n] != colors[(i-1+n)%n] {
			cnt++
		} else {
			cnt = 1
		}

		if cnt >= k {
			res++
		}
	}

	return res
}

func Test3208(t *testing.T) {
	colors := []int{0, 1, 0, 0, 1, 0, 1}
	k := 6
	res := numberOfAlternatingGroups1(colors, k)
	fmt.Println(res)
}
