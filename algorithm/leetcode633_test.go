package algorithm

import (
	"fmt"
	"math"
	"testing"
)

// 超时了
func judgeSquareSum(c int) bool {

	for i := 0; i*i <= c; i++ {
		for j := i; j*j <= c; j++ {
			if i*i+j*j == c {
				return true
			}
		}
	}

	return false
}

// 官方解法，双指针
func judgeSquareSum1(c int) bool {

	left := 0
	right := int(math.Sqrt(float64(c)))

	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum < c {
			left++
		} else {
			right--
		}
	}

	return false
}

func Test633(t *testing.T) {
	res := judgeSquareSum1(2)
	fmt.Println(res)
}
