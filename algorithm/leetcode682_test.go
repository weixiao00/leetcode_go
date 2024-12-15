package algorithm

import (
	"fmt"
	"strconv"
	"testing"
)

// 模拟
func calPoints(operations []string) int {

	if len(operations) == 0 {
		return 0
	}

	res := 0
	score := make([]int, 0)
	for _, operation := range operations {
		if operation == "+" {
			score = append(score, score[len(score)-1]+score[len(score)-2])
		} else if operation == "C" {
			score = score[0 : len(score)-1]
		} else if operation == "D" {
			score = append(score, 2*score[len(score)-1])
		} else {
			integer, _ := strconv.Atoi(operation)
			score = append(score, integer)
		}
	}

	for i := range score {
		res += score[i]
	}

	return res
}

func Test682(t *testing.T) {
	ops := []string{"5", "2", "C", "D", "+"}
	res := calPoints(ops)
	fmt.Println(res)
}
