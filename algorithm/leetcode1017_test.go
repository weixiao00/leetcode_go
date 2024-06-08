package algorithm

import (
	"fmt"
	"strconv"
	"testing"
)

// 进制转换
func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}

	bit := [32]int{}
	// 计算进位
	for i := 0; i < 32 && n != 0; i++ {
		if n&1 == 1 {
			// 下边为奇数
			bit[i] += 1
			if i&1 == 1 {
				bit[i+1] += 1
			}
		}
		n = n >> 1
	}

	carry := 0
	for i := 0; i < 32; i++ {
		val := carry + bit[i]
		bit[i] = val & 1
		carry = (val - bit[i]) / -2
	}

	i := 31
	for bit[i] == 0 {
		i--
	}

	var s string
	for i >= 0 {
		s += strconv.Itoa(bit[i])
		i--
	}

	return s
}

func Test1017(t *testing.T) {
	neg2 := baseNeg2(4)
	fmt.Println(neg2)
	//fmt.Println((-2) ^ 2)
}
