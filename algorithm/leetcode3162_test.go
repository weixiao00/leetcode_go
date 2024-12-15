package algorithm

import (
	"fmt"
	"testing"
)

func numberOfPairs(nums1 []int, nums2 []int, k int) int {

	mp1 := make(map[int]int)
	mp2 := make(map[int]int)

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	max1 := 0
	for _, num := range nums1 {
		mp1[num] = mp1[num] + 1
		max1 = getMax(max1, num)
	}

	for _, num := range nums2 {
		mp2[num] = mp2[num] + 1
	}

	res := 0
	for a, count2 := range mp2 {
		for b := a * k; b <= max1; b += a * k {
			count1, ok := mp1[b]
			if ok {
				res += count1 * count2
			}
		}
	}

	return res
}

func Test3162(t *testing.T) {
	nums1 := []int{1, 10, 11}
	nums2 := []int{2, 2, 2}
	k := 5
	res := numberOfPairs(nums1, nums2, k)
	fmt.Println(res)
}
