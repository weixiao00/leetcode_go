package algorithm

import (
	"fmt"
	"testing"
)

// 暴力
func largestTimeFromDigits(arr []int) string {

	if len(arr) == 0 {
		return ""
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := -1
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			for k := 0; k < len(arr); k++ {
				if k == i || k == j {
					continue
				}
				idx := 6 - i - j - k
				hour := arr[i]*10 + arr[j]
				min := arr[k]*10 + arr[idx]
				if hour >= 24 || min >= 60 {
					continue
				}
				res = getMax(res, hour*60+min)
			}
		}
	}

	if res == -1 {
		return ""
	}

	return fmt.Sprintf("%.2d:%.2d", res/60, res%60)
}

func Test949(t *testing.T) {
	arr := []int{0, 0, 0, 0}
	res := largestTimeFromDigits(arr)
	fmt.Println(res)
}
