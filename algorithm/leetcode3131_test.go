package algorithm

import (
	"fmt"
	"testing"
)

// 官方题解，最大的数加一个数依然是最大的数
func addedInteger1(nums1 []int, nums2 []int) int {

	max1 := 0
	for _, num1 := range nums1 {
		if num1 > max1 {
			max1 = num1
		}
	}

	max2 := 0
	for _, num2 := range nums2 {
		if num2 > max2 {
			max2 = num2
		}
	}

	return max2 - max1
}

// 循环两次
func addedInteger(nums1 []int, nums2 []int) int {

	numCountMap := make(map[int]int)

	nums1 = unique(nums1)
	nums2 = unique(nums2)
	for _, num2 := range nums2 {
		for _, num1 := range nums1 {
			val, ok := numCountMap[num2-num1]
			if !ok {
				numCountMap[num2-num1] = 1
			} else {
				numCountMap[num2-num1] = val + 1
			}
		}
	}

	for key, val := range numCountMap {
		if val == len(nums1) {
			return key
		}
	}
	return 0
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Test3131(t *testing.T) {
	nums1 := []int{7, 0, 7, 6, 7, 0, 9, 1}
	nums2 := []int{949, 955, 954, 948, 955, 948, 957, 955}
	res := addedInteger(nums1, nums2)
	fmt.Println(res)
}
