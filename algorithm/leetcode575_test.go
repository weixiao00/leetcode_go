package algorithm

import (
	"fmt"
	"testing"
)

func distributeCandies1(candyType []int) int {
	if len(candyType) == 0 {
		return 0
	}

	typeMap := make(map[int]bool, 0)

	for _, val := range candyType {
		_, ok := typeMap[val]
		if !ok {
			typeMap[val] = true
		}
	}

	return getMin2(len(candyType)/2, len(typeMap))
}

func getMin2(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Test575(t *testing.T) {
	candyType := []int{1, 1, 2, 2, 3, 3}
	candies := distributeCandies1(candyType)
	fmt.Println(candies)
}
