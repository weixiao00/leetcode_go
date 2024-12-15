package algorithm

import (
	"fmt"
	"testing"
)

// 快排思想
// 没过？？？
// java实现了
func findKthLargest(nums []int, k int) int {

	if len(nums) < k {
		return 0
	}

	var partition func(nums []int, begin, end int) int
	partition = func(nums []int, low, high int) int {
		x := nums[low]
		for low < high {
			for low < high && nums[high] <= x {
				high--
			}
			if low < high {
				nums[low] = nums[high]
				low++
			}
			for low < high && nums[low] >= x {
				low++
			}
			if low < high {
				nums[high] = nums[low]
				high--
			}
		}

		nums[low] = x

		return low
	}

	var quickSelect func(nums []int, left, right, k int) int
	quickSelect = func(nums []int, left, right, k int) int {
		pivot := partition(nums, left, right)
		if pivot == k {
			return nums[pivot]
		} else if pivot < k {
			return quickSelect(nums, pivot+1, right, k)
		} else {
			return quickSelect(nums, left, pivot-1, k)
		}
	}

	// 这个就不过，会超时
	//pos := partition(nums, 0, len(nums)-1)
	//for pos != k-1 {
	//	if pos < k-1 {
	//		pos = partition(nums, pos+1, len(nums)-1)
	//	} else {
	//		pos = partition(nums, 0, pos-1)
	//	}
	//}

	return quickSelect(nums, 0, len(nums)-1, k-1)
}

//func findKthLargest1(nums []int, k int) int {
//	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
//	heap.Init(nums)
//}

func Test215(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	res := findKthLargest(nums, k)
	fmt.Println(res)
}
