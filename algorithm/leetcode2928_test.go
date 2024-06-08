package algorithm

import (
	"fmt"
	"testing"
)

func distributeCandies(n int, limit int) int {

	res := 0
	for i := 0; i <= limit; i++ {
		idle := n - i
		if idle > 2*limit {
			continue
		}
		if idle < 0 {
			continue
		}

		// 第二个小朋友可以分到的最大值
		maxVal := getMin3(limit, idle)
		// 第二个小朋友可以分到的最小值
		minVal := getMax8(0, idle-limit)

		res += maxVal - minVal + 1
	}

	return res
}

func getMax8(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin3(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test2928(t *testing.T) {
	n := 1
	limit := 3
	candies := distributeCandies(n, limit)
	fmt.Println(candies)
}
