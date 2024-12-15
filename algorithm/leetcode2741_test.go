package algorithm

import (
	"fmt"
	"testing"
)

const MOD = 1000000007

var f [][]int
var numsc []int
var n int

// 数组所有符合相邻元素除数为0的组成的数组个数
func specialPerm(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	res := 0
	numsc = nums

	// 用来存储循环到的位置。如果循环到直接返回
	f = make([][]int, 0)
	n = len(nums)
	fn := 1 << n
	for i := 0; i < fn; i++ {
		arr := make([]int, n)
		for j := range arr {
			arr[j] = -1
		}
		f = append(f, arr)
	}

	// 状态压缩，每一个数作为最后一个数
	for i := range nums {
		res = (res + dfs3(fn-1, i)) % MOD
	}

	return res
}

// i作为最后一个数
func dfs3(state, i int) int {
	// 已经循环过了
	if f[state][i] != -1 {
		return f[state][i]
	}
	// 最后一个数
	if state == (1 << i) {
		return 1
	}
	// 标记循环过
	f[state][i] = 0
	for j := 0; j < n; j++ {
		if i == j || state>>j&1 == 0 {
			continue
		}
		if numsc[i]%numsc[j] != 0 && numsc[j]%numsc[i] != 0 {
			continue
		}
		// 递归j作为最后一个数。将i位置置为0
		f[state][i] = (f[state][i] + dfs3(state^(1<<i), j)) % MOD
	}

	return f[state][i]
}

func Test2741(t *testing.T) {
	nums := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}
	res := specialPerm(nums)
	fmt.Println(res)
}
