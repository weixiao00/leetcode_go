package algorithm

import (
	"fmt"
	"testing"
)

func wateringPlants(plants []int, capacity int) int {
	if len(plants) == 0 {
		return 0
	}

	idx := -1
	res := 0
	residue := capacity
	for i, water := range plants {
		res += i - idx
		idx = i
		residue -= water
		if i == len(plants)-1 {
			break
		}
		if residue < plants[i+1] {
			idx = -1
			residue = capacity
			res += i - idx
		}
	}
	return res
}

func Test2079(t *testing.T) {
	plants := []int{1, 1, 1, 4, 2, 3}
	capacity := 4
	res := wateringPlants(plants, capacity)
	fmt.Println(res)
}
