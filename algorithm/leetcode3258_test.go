package algorithm

import (
	"fmt"
	"testing"
)

// 枚举
func countKConstraintSubstrings(s string, k int) int {

	res := 0
	for i := 0; i < len(s); i++ {
		count := [2]int{}
		for j := i; j < len(s); j++ {
			count[s[j]-'0']++
			if count[0] > k && count[1] > k {
				break
			}
			res++
		}
	}

	return res
}

func Test3258(t *testing.T) {
	s := "10101"
	k := 1
	res := countKConstraintSubstrings(s, k)
	fmt.Println(res)
}
