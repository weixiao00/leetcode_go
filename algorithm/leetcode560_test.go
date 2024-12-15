package algorithm

import (
	"fmt"
	"testing"
)

// 滑动窗口不行，因为可能有负数

//暴力解法O(n*n)
func subarraySum(nums []int, k int) int {

	if len(nums) == 0 {
		return 0
	}
	res := 0
	for begin := range nums {
		count := 0
		for end := begin; end < len(nums); end++ {
			count += nums[end]
			if count == k {
				res++
			}
		}
	}

	return res
}

// 前缀和+hash优化
// 和长度大于等于2的连续子数组%k==0的个数解法基本相同
// O(n)
// 前缀和一般是处理连续子数组的问题
func subarraySum1(nums []int, k int) int {

	if len(nums) == 0 {
		return 0
	}
	res := 0
	mp := make(map[int]int)
	count := 0
	mp[0] = 1
	for _, num := range nums {
		count += num
		val, ok := mp[count-k]
		if ok {
			res += val
		}
		// 可能有负数，所以是+1，上边是+val
		mp[count] += 1
	}

	return res
}

func Test560(t *testing.T) {
	nums := []int{1, 2, 3}
	res := subarraySum1(nums, 3)
	fmt.Println(res)
}
