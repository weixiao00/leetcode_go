package algorithm

import (
	"fmt"
	"testing"
)

func numJewelsInStones(jewels string, stones string) int {
	if len(jewels) == 0 || len(stones) == 0 {
		return 0
	}

	res := 0
	for _, s := range stones {
		for _, j := range jewels {
			if s == j {
				res++
			}
		}
	}

	return res
}

func Test711(t *testing.T) {
	jewels := "aA"
	stones := "aAAbbbb"
	inStones := numJewelsInStones(jewels, stones)
	fmt.Println(inStones)
}
