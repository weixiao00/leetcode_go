package algorithm

import "testing"

func Test2455(t *testing.T) {

}

func averageValue(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 0
	total := 0
	for _, num := range nums {
		if num%6 == 0 {
			total += num
			k++
		}
	}
	if k > 0 {
		return total / k
	}
	return 0
}
