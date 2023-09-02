package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func captureForts(forts []int) int {
	if forts == nil {
		return 0
	}

	res := 0
	pre := 0
	for i, fort := range forts {
		if fort == 1 || fort == -1 {
			if fort == -forts[pre] {
				res = int(math.Max(float64(res), float64(i-pre-1)))
			}

			pre = i
		}
	}

	return res
}

func Test2511(t *testing.T) {
	forts := []int{1, 0, 0, -1, 0, 0, 0, 0, 1}
	res := captureForts(forts)
	fmt.Println(res)
}
