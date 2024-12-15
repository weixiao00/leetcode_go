package algorithm

import (
	"fmt"
	"testing"
)

func findIntersectionValues(nums1 []int, nums2 []int) []int {

	mp1 := make(map[int]bool)
	mp2 := make(map[int]bool)

	for _, num := range nums1 {
		mp1[num] = true
	}
	for _, num := range nums2 {
		mp2[num] = true
	}

	res1 := 0
	for _, key := range nums1 {
		if mp2[key] {
			res1++
		}
	}
	res2 := 0
	for _, key := range nums2 {
		if mp1[key] {
			res2++
		}
	}

	return []int{res1, res2}
}

func Test2956(t *testing.T) {
	nums1 := []int{4, 3, 2, 3, 1}
	nums2 := []int{2, 2, 5, 2, 3, 6}
	res := findIntersectionValues(nums1, nums2)
	fmt.Println(res)
}
