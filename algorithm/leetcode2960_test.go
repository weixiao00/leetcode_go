package algorithm

import (
	"fmt"
	"testing"
)

func countTestedDevices(batteryPercentages []int) int {
	if len(batteryPercentages) == 0 {
		return 0
	}
	res := 0
	for _, b := range batteryPercentages {
		if b-res > 0 {
			res++
		}
	}

	return res
}

func Test2960(t *testing.T) {
	arr := []int{0, 1, 2}
	devices := countTestedDevices(arr)
	fmt.Println(devices)
}
