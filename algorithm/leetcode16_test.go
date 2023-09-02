package algorithm

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func threeSumClosest(nums []int, target int) int {
	best := math.MaxInt
	if nums == nil || len(nums) == 0 {
		return 0
	}
	length := len(nums)
	sort.Ints(nums)

	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	for k, num := range nums {
		i := k + 1
		j := length - 1
		if k > 0 && num == nums[k-1] {
			continue
		}
		for i < j {
			cur := num + nums[i] + nums[j]
			if cur == target {
				return target
			}

			//if abs(cur, target) < abs(best, target) {
			//	best = cur
			//}
			update(cur)
			if cur > target {
				j0 := j
				for j0 > i && nums[j] == nums[j0] {
					j0--
				}
				j = j0
			} else {
				i0 := i
				for i0 < j && nums[i0] == nums[i] {
					i0++
				}
				i = i0
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// 这样求绝对值容易值越界
//func abs(cur, target int) int {
//	if cur < target {
//		return target - cur
//	} else {
//		return cur - target
//	}
//}

func TestName(t *testing.T) {
	arr := make([]int, 0, 4)
	arr = append(arr, 1, 1, 1, 1)
	fmt.Println(threeSumClosest(arr, -100))
}
