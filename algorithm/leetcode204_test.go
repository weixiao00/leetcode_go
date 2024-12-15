package algorithm

import (
	"fmt"
	"testing"
)

// 埃氏筛(质数筛)
// 1和本身两个因数
func countPrimes(n int) int {

	res := 0
	filter := make([]bool, n)
	for i := range filter {
		filter[i] = true
	}
	for i := 2; i < n; i++ {
		if filter[i] {
			res++
		}
		// i是质数，2i，3i，4i......一定不是质数
		for j := i * 2; j < n; j += i {
			filter[j] = false
		}
	}

	return res
}

func Test204(t *testing.T) {
	res := countPrimes(10)
	fmt.Println(res)
}
