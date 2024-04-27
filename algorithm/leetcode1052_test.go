package algorithm

import (
	"fmt"
	"testing"
)

func maxSatisfied(customers []int, grumpy []int, minutes int) int {

	if len(customers) == 0 {
		return 0
	}

	total := 0
	for i, customerNum := range customers {
		if grumpy[i] == 0 {
			total += customerNum
		}
	}

	maxIncrease := 0
	for i := 0; i < minutes; i++ {
		maxIncrease += customers[i] * grumpy[i]
	}

	increase := maxIncrease
	for i := minutes; i < len(customers); i++ {
		increase += customers[i]*grumpy[i] - customers[i-minutes]*grumpy[i-minutes]
		maxIncrease = maxVal(maxIncrease, increase)
	}

	return total + maxIncrease
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test1052(t *testing.T) {
	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	minutes := 3
	satisfied := maxSatisfied(customers, grumpy, minutes)
	fmt.Println(satisfied)
}
