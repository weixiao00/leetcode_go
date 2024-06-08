package algorithm

import (
	"fmt"
	"strings"
	"testing"
)

/**
官方解法
*/
func garbageCollection1(garbage []string, travel []int) int {
	distance := make(map[rune]int)
	res := 0
	curDis := 0
	for i, item := range garbage {
		res += len(item)
		if i > 0 {
			curDis += travel[i-1]
		}
		for _, c := range item {
			distance[c] = curDis
		}
	}
	for _, v := range distance {
		res += v
	}
	return res
}

func garbageCollection(garbage []string, travel []int) int {
	if len(garbage) == 0 {
		return 0
	}
	mIdx := 0
	pIdx := 0
	gIdx := 0
	garbageNum := 0
	for i, gar := range garbage {
		garbageNum += len(gar)
		if strings.Contains(gar, "M") {
			mIdx = i
		}
		if strings.Contains(gar, "P") {
			pIdx = i
		}
		if strings.Contains(gar, "G") {
			gIdx = i
		}
	}

	res := garbageNum
	if mIdx > 0 {
		for i := 0; i < mIdx; i++ {
			res += travel[i]
		}
	}

	if pIdx > 0 {
		for i := 0; i < pIdx; i++ {
			res += travel[i]
		}
	}

	if gIdx > 0 {
		for i := 0; i < gIdx; i++ {
			res += travel[i]
		}
	}

	return res
}

func Test2391(t *testing.T) {
	garbage := []string{"MMM", "PGM", "GP"}
	travel := []int{3, 10}
	collection := garbageCollection(garbage, travel)
	fmt.Println(collection)
}
