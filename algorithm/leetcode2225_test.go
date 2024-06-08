package algorithm

import (
	"fmt"
	"sort"
	"testing"
)

func findWinners(matches [][]int) [][]int {
	if len(matches) == 0 {
		return matches
	}
	timeMap := make(map[int]int, 0)
	for _, winArr := range matches {
		win := winArr[0]
		loser := winArr[1]
		_, ok := timeMap[win]
		if !ok {
			timeMap[win] = 0
		}
		loserTime, ok := timeMap[loser]
		if !ok {
			timeMap[loser] = 1
		} else {
			timeMap[loser] = loserTime + 1
		}
	}

	answer := make([][]int, 2)
	for loser, loserTime := range timeMap {
		if loserTime > 1 {
			continue
		}
		axr := answer[loserTime]
		axr = append(axr, loser)
		answer[loserTime] = axr
	}

	// 排序
	for _, a := range answer {
		sort.Ints(a)
	}

	return answer
}

func Test2225(t *testing.T) {
	matches := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}
	winners := findWinners(matches)
	fmt.Println(winners)
}
