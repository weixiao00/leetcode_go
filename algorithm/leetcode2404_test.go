package algorithm

import (
	"fmt"
	"testing"
)

//[0,1,2,2,4,4,1]
func Test2404(t *testing.T) {
	// 切片定义，类似于java里的List集合
	//arr := make([]int, 5)
	//arr = append(arr, 1)
	// 数组定义，类似于java里的数组
	//var arr = []int{0,1,2,2,4,4,1}
	arr := []int{0, 1, 2, 2, 4, 4, 1}
	res := mostFrequentEven(arr)
	fmt.Println(res)
}

func mostFrequentEven(nums []int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	count := map[int]int{}
	for _, value := range nums {
		if value%2 == 0 {
			count[value]++
		}
	}

	fmt.Println(count)
	res, curCount := -1, 0
	for key, value := range count {
		if value > curCount || (value == curCount && key < res) {
			res = key
			curCount = value
		}
	}
	return res
}
