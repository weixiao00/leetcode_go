package algorithm

import (
	"fmt"
	"testing"
)

// 合并两个有序数组
// 1. 2合到1上，排序
// 2. 创建一个临时数组
// 3. 从后向前双指针
func merge1(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}

	p1 := m - 1
	p2 := n - 1

	i := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		cur := 0
		if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[i] = cur
		i--
	}
	// 只有num2有数据才会放到num1
	// num1有数据直接返回就行了。因为直接就在num1上
	if p1 < 0 {
		for p2 >= 0 {
			nums1[i] = nums2[p2]
			p2--
			i--
		}
	}
}

func Test88(t *testing.T) {
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	merge1(nums1, m, nums2, n)
	fmt.Println(nums1)
}
