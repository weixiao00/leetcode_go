package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	if newInterval == nil {
		return intervals
	}
	left := newInterval[0]
	right := newInterval[1]
	len := len(intervals)
	res := make([][]int, 0)
	placed := false
	for i := 0; i < len; i++ {
		interval := intervals[i]
		if interval[1] < left {
			res = append(res, interval)
		} else if interval[0] > right {
			if !placed {
				res = append(res, []int{left, right})
				placed = true
			}
			res = append(res, interval)
		} else {
			left = int(math.Min(float64(left), float64(interval[0])))
			right = int(math.Max(float64(right), float64(interval[1])))
		}
	}

	if !placed {
		res = append(res, []int{left, right})
	}

	return res
}

func Test57(t *testing.T) {
	//intervals = [[1,3],[6,9]], newInterval = [2,5]
	arr1 := []int{1, 3}
	arr2 := []int{6, 9}
	arr3 := []int{2, 5}
	arr := make([][]int, 0, 2)
	arr = append(arr, arr1)
	arr = append(arr, arr2)
	res := insert(arr, arr3)
	fmt.Println(res)
}
