package algorithm

import (
	"fmt"
	"math"
	"testing"
)

// 二分法
func repairCars(ranks []int, cars int) int64 {
	if len(ranks) == 0 {
		return 0
	}

	// 闭包可以用外部的参数。要不然用定义方法还得传参数进去。和java的函数式编程差不多
	check := func(time int) bool {
		carNum := 0
		for i, _ := range ranks {
			carNum += int(math.Sqrt(float64(time / ranks[i])))
		}

		return carNum >= cars
	}

	left := 1
	right := ranks[0] * cars * cars
	for left < right {
		middle := (left + right) >> 1
		if check(middle) {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return int64(left)
}

func check(rank []int, time int, cars int) bool {
	carNum := 0
	for i, _ := range rank {
		carNum += int(math.Sqrt(float64(time / rank[i])))
	}

	return carNum >= cars
}

func Test2594(t *testing.T) {
	arr := []int{4, 2, 3, 1}
	cars := repairCars(arr, 10)
	fmt.Println(cars)
}
