package algorithm

import (
	"fmt"
	"testing"
)

func longestEqualSubarray(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	// key:nums里的元素，value:nums里元素的下标
	timeMap := make(map[int][]int, 0)

	for i, val := range nums {
		timeArr, ok := timeMap[val]
		if !ok {
			timeArr = make([]int, 0)
			timeArr = append(timeArr, i)
			timeMap[val] = timeArr
			continue
		}
		timeArr = append(timeArr, i)
		timeMap[val] = timeArr
	}

	res := 0
	for _, timeArr := range timeMap {
		i := 0
		for j := range timeArr {
			// 总元素个数 - 相同的元素个数 = 不同元素个数
			for (timeArr[j]-timeArr[i]+1)-(j-i+1) > k {
				i++
			}
			res = getMax6(res, j-i+1)
		}
	}

	return res
}

// 官方题解2
func longestEqualSubarray1(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	i := 0
	timeMap := make(map[int]int, 0)
	for j, num := range nums {
		time, ok := timeMap[num]
		if !ok {
			timeMap[num] = 1
		} else {
			timeMap[num] = time + 1
		}
		// 滑动窗口变小
		// 总数-相同的个数
		for j-i+1-timeMap[nums[i]] > k {
			timeMap[nums[i]] = timeMap[nums[i]] - 1
			i++
		}

		res = getMax6(res, timeMap[num])
	}
	return res
}

func getMax6(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Test2831(t *testing.T) {
	nums := []int{1, 2, 1}
	k := 0
	subarray := longestEqualSubarray1(nums, k)
	fmt.Println(subarray)
}
