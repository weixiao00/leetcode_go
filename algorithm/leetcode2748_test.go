package algorithm

import (
	"fmt"
	"testing"
)

func countBeautifulPairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var getMax func(a, b int) int
	getMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var prime func(a, b int) int
	prime = func(a, b int) int {
		max := getMax(a, b)
		for i := 2; i <= max; i++ {
			if a%i == 0 && b%i == 0 {
				return i
			}
		}
		return 1
	}

	res := 0
	for i := range nums {
		for nums[i] >= 10 {
			nums[i] = nums[i] / 10
		}
		for j := i + 1; j < len(nums); j++ {
			if prime(nums[i], nums[j]%10) == 1 {
				res++
			}
		}
	}

	return res
}

//func gcd(a, b int) int {
//	for b != 0 {
//		a, b = b, a%b
//	}
//	return a
//}

func Test2748(t *testing.T) {
	nums := []int{756, 1324, 2419, 495, 106, 111, 1649, 1474, 2001, 1633, 273, 1804, 2102, 1782, 705, 1529, 1761, 1613, 111, 186, 412}
	res := countBeautifulPairs(nums)
	fmt.Println(res)
}
