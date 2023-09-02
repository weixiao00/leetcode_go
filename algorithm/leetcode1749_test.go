package algorithm

import (
	"fmt"
	"testing"
)

func maxAbsoluteSum(nums []int) int {
	if nums == nil {
		return 0
	}
	positiveMax := 0
	negativeMin := 0
	positiveCur := 0
	negativeCur := 0
	for _, num := range nums {
		positiveCur += num
		if positiveMax < positiveCur {
			positiveMax = positiveCur
		}
		if positiveCur <= 0 {
			positiveCur = 0
		}
		negativeCur += num
		if negativeCur < negativeMin {
			negativeMin = negativeCur
		}
		if negativeCur >= 0 {
			negativeCur = 0
		}
	}
	if positiveMax > -negativeMin {
		return positiveMax
	} else {
		return -negativeMin
	}
}

func Test1749(t *testing.T) {
	arr := []int{1, -3, 2, 3, -4}
	sum := maxAbsoluteSum(arr)
	fmt.Println(sum)
}
