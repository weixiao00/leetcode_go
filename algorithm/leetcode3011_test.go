package algorithm

import (
	"fmt"
	"math/bits"
	"testing"
)

// 分组计算，升序
func canSortArray(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	oneCount := 0
	lastGroupMax := 0
	curGroupMax := 0
	for _, num := range nums {

		// 一个组的
		if bits.OnesCount(uint(num)) == oneCount {
			curGroupMax = getMax(curGroupMax, num)
		} else {
			lastGroupMax = curGroupMax
			curGroupMax = num
			oneCount = bits.OnesCount(uint(num))
		}

		if lastGroupMax > num {
			return false
		}
	}

	return true
}

func Test3011(t *testing.T) {
	nums := []int{8, 4, 2, 30, 15}
	res := canSortArray(nums)
	fmt.Println(res)
}
