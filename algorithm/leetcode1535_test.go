package algorithm

import (
	"fmt"
	"testing"
)

func getWinner(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}

	prev := getMax4(arr[0], arr[1])
	maxNum := prev
	count := 1
	if k == 1 {
		return prev
	}
	for i := 2; i < len(arr); i++ {
		if prev > arr[i] {
			count++
			if count == k {
				return prev
			}
		} else {
			prev = arr[i]
			count = 1
		}
		maxNum = getMax4(maxNum, prev)
	}

	return maxNum
}

func getMax4(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test1535(t *testing.T) {
	arr := []int{2, 1, 3, 5, 4, 6, 7}
	k := 2
	winner := getWinner(arr, k)
	fmt.Println(winner)
}
