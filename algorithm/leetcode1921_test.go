package algorithm

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func eliminateMaximum(dist []int, speed []int) int {
	if dist == nil || speed == nil {
		return 0
	}
	arrives := make([]int, 0, len(dist))
	for i, _ := range dist {
		arrive := (dist[i]-1)/speed[i] + 1
		arrives = append(arrives, arrive)
	}

	sort.Ints(arrives)

	for i, _ := range arrives {
		if arrives[i] <= i {
			return i
		}
	}

	return len(dist)
}

//以下是自己的思路。超出时间限制了
func eliminateMaximumBack(dist []int, speed []int) int {
	if dist == nil || speed == nil {
		return 0
	}

	res := 0
	size := len(dist)
	i := 0
	for i < size {
		if eliminate(dist) {
			return res
		}
		// 开抢
		res++
		// 打死移除
		dist, speed = remove(dist, speed)
		advance(dist, speed)
		i++
	}
	return res
}

func eliminate(arr []int) bool {
	for _, val := range arr {
		if val <= 0 {
			return true
		}
	}

	return false
}

func advance(arr, speed []int) {
	size := len(arr)
	i := 0
	for i < size {
		arr[i] -= speed[i]
		i++
	}
}

func remove(arr, speed []int) ([]int, []int) {
	minIndex := 0
	minVal := math.MaxInt
	for i, val := range arr {
		if (val-1)/speed[i]+1 < minVal {
			minIndex = i
			minVal = (val-1)/speed[i] + 1
		}
	}

	arr1 := make([]int, 0, len(arr)-1)
	for i, val := range arr {
		if i != minIndex {
			arr1 = append(arr1, val)
		}
	}

	speed1 := make([]int, 0, len(arr)-1)
	for i, val := range speed {
		if i != minIndex {
			speed1 = append(speed1, val)
		}
	}

	return arr1, speed1
}

func Test1921(t *testing.T) {
	dist := []int{3, 3, 5, 7, 7}
	speed := []int{1, 1, 4, 2, 2}
	res := eliminateMaximumBack(dist, speed)
	fmt.Println(res)
}

//------------------------------------------------------------------------------

func advance1(arr, speed []*int) {
	size := len(arr)
	i := 0
	for i < size {
		*arr[i] -= *speed[i]
		i++
	}
}

// 输出r的值
func a(r *int) {
	res := 100
	*r = res
}

// 输出100
func b(r *int) {
	res := 100
	r = &res
}

func Test19211(t *testing.T) {
	//arr := []int{1, 3, 6}
	//speed := []int{1, 3, 5}
	//advance1(arr, speed)
	//for _, val := range arr {
	//	fmt.Println(val)
	//}
	res := 11
	a(&res)
	fmt.Println(res)
}
