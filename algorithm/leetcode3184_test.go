package algorithm

import (
	"fmt"
	"testing"
)

func countCompleteDayPairs(hours []int) int {

	if len(hours) == 0 {
		return 0
	}

	res := 0
	for i := range hours {
		for j := i + 1; j < len(hours); j++ {
			if (hours[i]+hours[j])%24 == 0 {
				res++
			}
		}
	}

	return res
}

func Test3184(t *testing.T) {
	hours := []int{12, 12, 30, 24, 24}
	res := countCompleteDayPairs(hours)
	fmt.Println(res)
}
