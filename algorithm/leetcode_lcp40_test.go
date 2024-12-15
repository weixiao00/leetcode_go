package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

// 贪心，奇偶性
func maxmiumScore(cards []int, cnt int) int {

	if len(cards) == 0 {
		return 0
	}

	sort.Ints(cards)

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	// 偶数
	odd := -1
	// 奇数
	even := -1
	count := 0
	i := len(cards) - 1
	for ; i >= 0; i-- {
		if cnt == 0 {
			break
		}
		cnt--
		card := cards[i]
		if card&1 == 0 {
			odd = card
		} else {
			even = card
		}
		count += card
	}

	if count&1 == 0 {
		return count
	}

	for ; i >= 0; i-- {
		if cards[i]&1 == 0 {
			if even != -1 {
				res = getMax(res, count-even+cards[i])
			}
		} else {
			if odd != -1 {
				res = getMax(res, count-odd+cards[i])
			}
		}
	}

	return res
}

func TestLcp40(t *testing.T) {
	cards := []int{1, 2, 8, 9}
	cnt := 3
	res := maxmiumScore(cards, cnt)
	fmt.Println(res)
}
