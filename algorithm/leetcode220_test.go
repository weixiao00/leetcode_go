package algorithm

import (
	"fmt"
	"testing"
)

// 三种解法
// 滑动窗口，计算valueDiff窗口是否满足，O(nk)
// 滑动窗口，k数组排序，二叉树，O(nlogk)
// 滑动窗口，桶排序排序，O(n)
// 我这里实现第三种桶排序
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {

	if len(nums) == 0 {
		return true
	}

	var getID func(num, weight int) int
	getID = func(num, weight int) int {
		if num >= 0 {
			return num / weight
		}
		return num/weight - 1
	}

	var abs func(a, b int) int
	abs = func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	mp := make(map[int]int)
	for i, num := range nums {
		weight := valueDiff + 1
		id := getID(num, weight)
		_, ok := mp[id]
		if ok {
			return true
		}
		// 相邻的
		val1, ok := mp[id+1]
		if ok && abs(val1, num) <= valueDiff {
			return true
		}
		// 相邻的
		val0, ok := mp[id-1]
		if ok && abs(val0, num) <= valueDiff {
			return true
		}
		// 放入map
		mp[id] = num
		// 窗口大于，删除最远的
		if i >= indexDiff {
			delete(mp, getID(nums[i-indexDiff], weight))
		}
	}

	return false
}

func Test220(t *testing.T) {
	nums := []int{1, 5, 9, 1, 5, 9}
	indexDiff := 2
	valueDiff := 3
	res := containsNearbyAlmostDuplicate(nums, indexDiff, valueDiff)
	fmt.Println(res)
}
