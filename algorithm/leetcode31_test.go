package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

// 下一个排列
func nextPermutation(nums []int) {

	if len(nums) == 0 {
		return
	}

	p := -1
	// 找到最后一个小于的
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			p = i
		}
	}

	if p == -1 {
		sort.Ints(nums)
		return
	}

	q := -1
	// p之后找到大于nums[p]，最远的
	for j := p + 1; j < len(nums); j++ {
		if nums[p] < nums[j] {
			q = j
		}
	}

	// 交换，排序
	tmp := nums[p]
	nums[p] = nums[q]
	nums[q] = tmp

	i := p + 1
	j := len(nums) - 1

	for i < j {
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
		i++
		j--
	}
}

func Test31(t *testing.T) {
	nums := []int{1, 2, 3, 5, 4}
	nextPermutation(nums)
	fmt.Println(nums)
}
