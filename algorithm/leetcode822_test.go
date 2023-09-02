package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func flipgame(fronts []int, backs []int) int {
	if fronts == nil || backs == nil {
		return 0
	}

	sameMap := map[int]interface{}{}
	size := len(fronts)
	for i := 0; i < size; i++ {
		if fronts[i] == backs[i] {
			sameMap[fronts[i]] = struct{}{}
		}
	}
	res := math.MaxInt
	for _, front := range fronts {
		_, ok := sameMap[front]
		if front < res && !ok {
			res = front
		}
	}

	for _, back := range backs {
		_, ok := sameMap[back]
		if back < res && !ok {
			res = back
		}
	}

	if res == math.MaxInt {
		return 0
	}
	return res
}

func same(fronts []int, front, back int, res *int) int {
	if front == back {
		return math.MaxInt
	}
	flag := true
	for _, f := range fronts {
		if f == back {
			flag = false
			continue
		}
	}
	if flag {
		return back
	}

	return math.MaxInt
}

func Test822(t *testing.T) {
	fronts := []int{2, 1}
	backs := []int{1, 2}
	i := flipgame(fronts, backs)
	fmt.Println(i)
}
