package algorithm

import (
	"fmt"
	"testing"
)

func sumOfTheDigitsOfHarshadNumber(x int) int {

	if x == 0 {
		return 0
	}

	y := x
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}

	if y%sum == 0 {
		return sum
	}

	return -1
}

func Test3099(t *testing.T) {
	res := sumOfTheDigitsOfHarshadNumber(18)
	fmt.Println(res)
}
