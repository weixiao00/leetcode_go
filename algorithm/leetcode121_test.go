package algorithm

import "math"

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var getMin func(a, b int) int
	getMin = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	min := math.MaxInt32
	res := 0
	for _, price := range prices {
		min = getMin(min, price)
		profit := price - min
		res = getMax(res, profit)
	}

	return res
}
