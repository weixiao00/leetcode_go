package algorithm

import (
	"fmt"
	"testing"
)

func temperatureTrend(temperatureA []int, temperatureB []int) int {

	var getLetter func(a, b int) int
	getLetter = func(a, b int) int {
		if a == b {
			return 0
		}
		if a > b {
			return -1
		}
		return 1
	}

	len := len(temperatureA)

	res := 0
	cur := 0
	for i := 0; i < len-1; i++ {
		letter1 := getLetter(temperatureA[i], temperatureA[i+1])
		letter2 := getLetter(temperatureB[i], temperatureB[i+1])
		if letter1 == letter2 {
			cur++
			res = getMax18(res, cur)
			continue
		}
		cur = 0
	}

	return res
}

func getMax18(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestLcp61(t *testing.T) {
	temperatureA := []int{21, 18, 18, 18, 31}
	temperatureB := []int{34, 32, 16, 16, 17}
	res := temperatureTrend(temperatureA, temperatureB)
	fmt.Println(res)
}
