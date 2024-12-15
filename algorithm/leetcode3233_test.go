package algorithm

import (
	"fmt"
	"math"
	"testing"
)

// 第二个因数是质数
// 特殊数字是平方数
func nonSpecialCount(l int, r int) int {

	res := r - l + 1
	n := int(math.Sqrt(float64(r)))
	filter := make([]bool, n+1)
	for i := range filter {
		filter[i] = true
	}
	for i := 2; i <= n; i++ {
		if filter[i] {
			// 是一个特殊数字
			if i*i >= l && i*i <= r {
				res--
			}
			// 标记非质数
			for j := 2 * i; j <= n; j += i {
				filter[j] = false
			}
		}
	}

	return res
}

func Test3233(t *testing.T) {
	res := nonSpecialCount(4, 16)
	fmt.Println(res)
}
