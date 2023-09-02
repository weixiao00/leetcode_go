package algorithm

import (
	"fmt"
	"testing"
)

func subtractProductAndSum(n int) int {

	sum := 0
	multi := 1

	for n > 0 {
		sum += n % 10
		multi *= n % 10
		n /= 10
	}

	return multi - sum
}

func Test1281(t *testing.T) {
	sum := subtractProductAndSum(4421)
	fmt.Println(sum)
}
