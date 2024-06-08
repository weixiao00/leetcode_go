package algorithm

import (
	"fmt"
	"testing"
)

func distributeCandies2(candies int, num_people int) []int {

	res := make([]int, num_people)
	idx := 0
	tmpCandies := 0
	for candies > 0 {
		if idx == num_people {
			idx = 0
		}
		tmpCandies += 1
		if candies-tmpCandies > 0 {
			candies -= tmpCandies
		} else {
			tmpCandies = candies
			candies = 0
		}
		res[idx] += tmpCandies
		idx++
	}

	return res
}

func Test1103(t *testing.T) {
	candies := 10
	num_people := 3
	res := distributeCandies2(candies, num_people)
	fmt.Println(res)
}
