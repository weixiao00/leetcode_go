package algorithm

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	if len(horizontalCuts) == 0 || len(verticalCuts) == 0 {
		return 0
	}

	horizontalCuts = append(horizontalCuts, 0, h)
	verticalCuts = append(verticalCuts, 0, w)
	sort.Ints(horizontalCuts)
	height := 0
	for i := 1; i < len(horizontalCuts); i++ {
		pre := horizontalCuts[i-1]
		cur := horizontalCuts[i]
		height = int(math.Max(float64(height), float64(cur-pre)))
	}
	sort.Ints(verticalCuts)
	width := 0
	for i := 1; i < len(verticalCuts); i++ {
		pre := verticalCuts[i-1]
		cur := verticalCuts[i]
		width = int(math.Max(float64(width), float64(cur-pre)))
	}

	return (height * width) % (1000000007)
}

func Test1465(t *testing.T) {
	//h = 5, w = 4, horizontalCuts = [1,2,4], verticalCuts = [1,3]
	horizontalCuts := []int{3}
	verticalCuts := []int{3}
	area := maxArea(5, 4, horizontalCuts, verticalCuts)
	fmt.Println(area)
}
