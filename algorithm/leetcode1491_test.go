package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func average(salary []int) float64 {
	if len(salary) == 0 {
		return 0
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

	top := math.MinInt
	bottom := math.MaxInt
	res := 0
	for _, s := range salary {
		top = getMax(top, s)
		bottom = getMin(bottom, s)
		res += s
	}

	return float64(res-top-bottom) / float64(len(salary)-2)
}

func Test1491(t *testing.T) {
	salary := []int{6000, 5000, 4000, 3000, 2000, 1000}
	value := average(salary)
	fmt.Println(value)
}
