package algorithm

import (
	"fmt"
	"testing"
)

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var res = int64(total/cost2) + 1

	internalTotal := total
	for internalTotal >= cost1 {
		internalTotal -= cost1
		num := int64(internalTotal/cost2) + 1
		res += num
	}

	return res
}

func Test2240(t *testing.T) {
	pencils := waysToBuyPensPencils(5, 10, 10)
	fmt.Println(pencils)
}
