package algorithm

import (
	"fmt"
	"testing"
)

// 单调栈
func dailyTemperatures(temperatures []int) []int {

	if len(temperatures) == 0 {
		return []int{}
	}

	answer := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			j := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			answer[j] = i - j
		}
		stack = append(stack, i)
	}

	return answer
}

func Test739(t *testing.T) {
	//[1,1,4,2,1,1,0,0]
	//[1 1 1 1 1 1 1 0]
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans := dailyTemperatures(temperatures)
	fmt.Println(ans)
}
