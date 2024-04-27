package algorithm

import (
	"fmt"
	"testing"
)

func combinationSum4(nums []int, target int) int {

	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func Test337(t *testing.T) {
	arr := []int{1, 2, 3}
	sum4 := combinationSum4(arr, 4)
	fmt.Println(sum4)
}
