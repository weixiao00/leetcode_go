package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {

	if len(nums) == 0 {
		return [][]int{}
	}

	res := make([][]int, 0)
	//先排序
	sort.Ints(nums)
	idx := 0
	for idx < len(nums) {
		num := nums[idx]
		left := idx + 1
		right := len(nums) - 1

		for left < right {
			// 相等的三元组
			if num+nums[left]+nums[right] == 0 {
				res = append(res, []int{num, nums[left], nums[right]})
				left++
				// 因为返回的是没有重复的三元组
				for left < right && nums[left-1] == nums[left] {
					left++
				}
				right--
				for left < right && nums[right+1] == nums[right] {
					right--
				}
			} else if num+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
		// 固定的那个，也不能相等
		idx++
		for idx < len(nums) && nums[idx-1] == nums[idx] {
			idx++
		}
	}

	return res
}

func Test15(t *testing.T) {
	// -4,-1,-1,0,1,2
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}
