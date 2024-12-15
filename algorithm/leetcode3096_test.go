package algorithm

import (
	"fmt"
	"testing"
)

// 前缀和
func minimumLevels(possible []int) int {
	if len(possible) == 0 {
		return 0
	}

	total := 0
	for _, p := range possible {
		if p == 0 {
			total -= 1
		} else {
			total += 1
		}
	}

	totalA := 0
	if possible[0] == 1 {
		totalA = 1
	} else {
		totalA = -1
	}

	for i := 1; i < len(possible); i++ {
		if totalA > total-totalA {
			return i
		}
		if possible[i] == 1 {
			totalA += 1
		} else {
			totalA -= 1
		}
	}

	return -1
}

func Test3096(t *testing.T) {
	possible := []int{1, 1, 1, 1, 1}
	res := minimumLevels(possible)
	fmt.Println(res)
}
