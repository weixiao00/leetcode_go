package algorithm

import (
	"fmt"
	"testing"
)

// 符合条件的连续子数组
// 前缀和+prefixSum[i] - prefixSum[j]除k余数为0。那么prefixSum[i]和prefixSum[j]余数相同
func checkSubarraySum(nums []int, k int) bool {

	if len(nums) < 2 {
		return false
	}

	// key: 余数，value：索引
	mp := make(map[int]int)

	// 前缀和为0
	mp[0] = -1
	remainder := 0
	for i, num := range nums {
		remainder = (remainder + num) % k
		preIndex, ok := mp[remainder]
		if ok {
			if i-preIndex >= 2 {
				return true
			}
		} else {
			mp[remainder] = i
		}
	}

	return false
}

func Test523(t *testing.T) {
	nums := []int{23, 2, 4, 6, 7}
	k := 6
	res := checkSubarraySum(nums, k)
	fmt.Println(res)
}
