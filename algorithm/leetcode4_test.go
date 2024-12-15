package algorithm

import (
	"fmt"
	"testing"
)

// 二分查找
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l1 := len(nums1)
	l2 := len(nums2)
	total := l1 + l2

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// 寻找第k小的数
	var getEle func(k int) int
	getEle = func(k int) int {

		index1 := 0
		index2 := 0
		for true {
			// 边界
			if index1 == l1 {
				return nums2[index2+k-1]
			}
			if index2 == l2 {
				return nums1[index1+k-1]
			}
			if k == 1 {
				return getMin(nums1[index1], nums2[index2])
			}

			half := k / 2
			newIndex1 := getMin(index1+half, l1) - 1
			newIndex2 := getMin(index2+half, l2) - 1
			pivot1 := nums1[newIndex1]
			pivot2 := nums2[newIndex2]
			if pivot1 <= pivot2 {
				k -= newIndex1 - index1 + 1
				index1 = newIndex1 + 1
			} else {
				k -= newIndex2 - index2 + 1
				index2 = newIndex2 + 1
			}
		}

		return -1
	}

	k := total / 2
	// 奇数个
	if total&1 == 1 {
		return float64(getEle(k + 1))
	} else {
		return float64(getEle(k)+getEle(k+1)) / 2
	}
}

func Test4(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
}
