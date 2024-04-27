package algorithm

import (
	"fmt"
	"testing"
)

func distanceTraveled(mainTank int, additionalTank int) int {
	if mainTank <= 0 {
		return 0
	}

	total := 0
	for mainTank >= 5 {
		five := mainTank - 5
		total += 5 * 10
		mainTank = five
		if additionalTank > 0 {
			additionalTank--
			mainTank++
		}
	}
	total += mainTank * 10
	return total
}

func Test2739(t *testing.T) {
	traveled := distanceTraveled(9, 2)
	fmt.Println(traveled)
}
