package algorithm

import (
	"container/list"
	"fmt"
	"testing"
)

// 1. 优先队列。O(nlgn) 需要存储{nums[i], i}
// 2. 单调递减的队列。O(n)
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	deque := list.New()
	for i := 0; i < k; i++ {
		for deque.Len() != 0 && nums[deque.Back().Value.(int)] < nums[i] {
			deque.Remove(deque.Back())
		}
		deque.PushBack(i)
	}
	res := make([]int, 0, len(nums)-k+1)
	res = append(res, nums[deque.Front().Value.(int)])

	for i := k; i < len(nums); i++ {
		for deque.Len() != 0 && nums[deque.Back().Value.(int)] < nums[i] {
			deque.Remove(deque.Back())
		}
		deque.PushBack(i)
		// 超过范围的移除
		for deque.Front().Value.(int)+k <= i {
			deque.Remove(deque.Front())
		}
		res = append(res, nums[deque.Front().Value.(int)])
	}

	return res
}

func Test239(t *testing.T) {
	arr := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	window := maxSlidingWindow(arr, k)
	fmt.Println(window)
}
