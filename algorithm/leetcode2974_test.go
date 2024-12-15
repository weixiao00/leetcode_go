package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func numberGame(nums []int) []int {

	if len(nums) < 2 {
		return []int{}
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		nums[i] = nums[i] + nums[i+1]
		nums[i+1] = nums[i] - nums[i+1]
		nums[i] = nums[i] - nums[i+1]

		// a = a^b
		// b = a^b
		// a = a^b
	}

	return nums
}

func Test2974(t *testing.T) {
	nums := []int{5, 4, 2, 3}
	res := numberGame(nums)
	fmt.Println(res)
}
