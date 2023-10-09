package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func minNumber(nums1 []int, nums2 []int) int {

	// 判断有没有相等的数字
	res := same(nums1, nums2)
	if res != math.MaxInt {
		return res
	}
	// 没有的话取两个数组最小的数
	res1 := min(nums1)
	res2 := min(nums2)
	return int(math.Min(float64(res1*10+res2), float64(res2*10+res1)))
}

func same(num1, num2 []int) int {

	num1Map := make(map[int]struct{}, len(num1))
	for _, val := range num1 {
		num1Map[val] = struct{}{}
	}

	res := math.MaxInt
	for _, val := range num2 {
		_, ok := num1Map[val]
		if ok {
			res = int(math.Min(float64(res), float64(val)))
		}
	}

	return res
}

func min(arr []int) int {
	if len(arr) == 0 {
		panic("empty array")
	}
	minVal := arr[0]
	for _, val := range arr {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func Test2605(t *testing.T) {
	num1 := []int{4, 1, 3}
	num2 := []int{5, 7}
	number := minNumber(num1, num2)
	fmt.Println(number)
}

func Test26051(t *testing.T) {
	num1Map := make(map[int]bool)
	s := num1Map[8]
	fmt.Println(s)
}
