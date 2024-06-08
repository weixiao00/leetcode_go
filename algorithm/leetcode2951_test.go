package algorithm

import (
	"fmt"
	"testing"
)

func findPeaks(mountain []int) []int {

	res := make([]int, 0)
	if len(mountain) < 3 {
		return res
	}

	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			res = append(res, i)
		}
	}

	return res
}

func Test2951(t *testing.T) {
	mountain := []int{2, 4, 4}
	res := findPeaks(mountain)
	fmt.Println(res)
}
