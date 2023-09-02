package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func maxDistToClosest(seats []int) int {
	if seats == nil {
		return 0
	}

	first := -1
	last := -1
	n := len(seats)
	d := 0

	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			if last != -1 {
				d = int(math.Max(float64(d), float64((i-last)>>1)))
			}
			if first == -1 {
				first = i
			}
			last = i
		}
	}

	return int(math.Max(float64(d), math.Max(float64(first), float64(n-1-last))))
}

func Test849(t *testing.T) {
	arr := []int{1, 0, 0, 0, 1, 0, 1}
	closest := maxDistToClosest(arr)
	fmt.Println(closest)
}
