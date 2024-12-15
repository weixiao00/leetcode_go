package algorithm

import (
	"fmt"
	"testing"
)

// 模拟
func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {

	runes1 := []rune(coordinate1)
	runes2 := []rune(coordinate2)

	for i := range runes1 {
		if (runes1[i]-runes2[i])&1 == 0 && (runes1[i+1]-runes2[i+1])&1 == 0 {
			return true
		} else if (runes1[i]-runes2[i])&1 == 1 && (runes1[i+1]-runes2[i+1])&1 == 1 {
			return true
		} else {
			return false
		}
	}

	return false
}

func Test3274(t *testing.T) {
	coordinate1 := "h7"
	coordinate2 := "c8"
	res := checkTwoChessboards(coordinate1, coordinate2)
	fmt.Println(res)
}
