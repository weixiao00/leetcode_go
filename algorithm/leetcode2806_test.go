package algorithm

import (
	"fmt"
	"testing"
)

func accountBalanceAfterPurchase(purchaseAmount int) int {

	div := purchaseAmount / 10
	if abs1(div*10, purchaseAmount) < abs1(div*10+10, purchaseAmount) {
		return 100 - div*10
	}
	return 100 - div*10 - 10
}

func abs1(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Test2806(t *testing.T) {
	purchase := accountBalanceAfterPurchase(15)
	fmt.Println(purchase)
}
