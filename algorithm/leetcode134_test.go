package algorithm

import (
	"fmt"
	"testing"
)

// 贪心算法
// 走到最远，走不到跳过去
func canCompleteCircuit(gas []int, cost []int) int {

	if len(gas) == 0 {
		return 0
	}

	i := 0
	len := len(gas)
	for i < len {
		gasOfSum := 0
		costOfSum := 0
		cnt := 0
		for cnt < len {
			j := (cnt + i) % len
			gasOfSum += gas[j]
			costOfSum += cost[j]
			if gasOfSum < costOfSum {
				break
			}
			cnt++
		}
		if cnt == len {
			return i
		} else {
			i = i + 1 + cnt
		}
	}

	return -1
}

func Test134(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	res := canCompleteCircuit(gas, cost)
	fmt.Println(res)
}
