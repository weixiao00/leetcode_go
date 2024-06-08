package algorithm

import (
	"fmt"
	"testing"
)

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	if len(plants) == 0 {
		return 0
	}

	i := 0
	j := len(plants) - 1
	stockCapA := capacityA
	stockCapB := capacityB
	res := 0
	for i <= j {
		if i == j {
			if stockCapA < stockCapB {
				if stockCapB < plants[i] {
					res++
				}
				break
			} else {
				if stockCapA < plants[i] {
					res++
				}
				break
			}
		}

		if stockCapA < plants[i] && plants[i] <= capacityA {
			res++
			stockCapA = capacityA
		}

		if stockCapA >= plants[i] {
			stockCapA -= plants[i]
			i++
		}

		if stockCapB < plants[j] && plants[j] <= capacityB {
			res++
			stockCapB = capacityB
		}

		if stockCapB >= plants[j] {
			stockCapB -= plants[j]
			j--
		}
	}

	return res
}

func Test2105(t *testing.T) {
	plants := []int{2, 1, 1}
	capacityA := 2
	capacityB := 2
	refill := minimumRefill(plants, capacityA, capacityB)
	fmt.Println(refill)
}
