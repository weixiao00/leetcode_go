package algorithm

import (
	"fmt"
	"testing"
)

// 状态压缩
// 字符串都是数字类型
func longestAwesome(s string) int {

	statusBitMap := make(map[int]int, 0)
	statusBitMap[0] = -1
	sequence := 0
	res := 0
	for j, ss := range s {
		sequence ^= 1 << (ss - '0')

		// aabbcc。这种情况。t(0,i-1)和t(0,j)相等
		i, ok := statusBitMap[sequence]
		if ok {
			res = getMax5(res, j-i)
		} else {
			statusBitMap[sequence] = j
		}
		// abbcc。这种情况。t(0,i-1)和t(0,j)最多有一位不相等
		for k := 0; k < 10; k++ {
			i, ok := statusBitMap[sequence^(1<<k)]
			if ok {
				res = getMax5(res, j-i)
			}
		}
	}
	return res
}

func getMax5(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test1542(t *testing.T) {
	s := "3242415"
	awesome := longestAwesome(s)
	fmt.Println(awesome)
}
