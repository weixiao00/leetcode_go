package algorithm

import (
	"fmt"
	"testing"
)

// 模拟
func numberOfAlternatingGroups(colors []int) int {

	if len(colors) < 3 {
		return 0
	}

	res := 0
	for i := range colors {
		left := i - 1
		right := i + 1
		if i == 0 {
			left = len(colors) - 1
		}
		if i == len(colors)-1 {
			right = 0
		}
		if colors[i] == 0 && colors[left] == 1 && colors[right] == 1 {
			res++
		}
		if colors[i] == 1 && colors[left] == 0 && colors[right] == 0 {
			res++
		}
	}

	return res
}

func Test3206(t *testing.T) {
	colors := []int{0, 1, 0, 0, 1}
	res := numberOfAlternatingGroups(colors)
	fmt.Println(res)
}
