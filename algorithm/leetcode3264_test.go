package algorithm

import (
	"fmt"
	"testing"
)

// 方法一：模拟
// 方法二：可以用优先队列
func getFinalState(nums []int, k int, multiplier int) []int {

	for i := 0; i < k; i++ {
		minIndex := 0
		for j := 0; j < len(nums); j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}
		nums[minIndex] = nums[minIndex] * multiplier
	}

	return nums
}

func Test3264(t *testing.T) {
	nums := []int{2, 1, 3, 5, 6}
	k := 5
	multiplier := 2
	res := getFinalState(nums, k, multiplier)
	fmt.Println(res)
}
