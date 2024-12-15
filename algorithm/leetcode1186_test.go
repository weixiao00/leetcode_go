package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func maximumSum(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 子数组删除一个了的最大值
	dp0 := 0

	// 所有之和
	dp1 := arr[0]

	res := math.MinInt32
	for i := 1; i < len(arr); i++ {
		dp0 = getMax(getMax(dp0+arr[i], dp1), arr[i])
		dp1 = getMax(0, dp1) + arr[i]
		res = getMax(res, getMax(dp0, dp1))
	}

	return res
}

func Test1186(t *testing.T) {
	arr := []int{-2, 5, 0, -1, 2}
	res := maximumSum(arr)
	fmt.Println(res)
}
