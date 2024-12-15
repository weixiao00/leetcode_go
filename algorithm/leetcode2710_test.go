package algorithm

import (
	"fmt"
	"testing"
)

func removeTrailingZeros(num string) string {

	if len(num) == 0 {
		return num
	}

	idx := len(num)
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] == '0' {
			idx = i
			continue
		}
		break
	}

	return num[0:idx]
}

func Test2710(t *testing.T) {
	num := "51230100"
	res := removeTrailingZeros(num)
	fmt.Println(res)
}
