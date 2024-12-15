package algorithm

import (
	"fmt"
	"testing"
)

//I	1
//V	5
//X	10
//L	50
//C	100
//D	500
//M	1000
// 整数转罗马

func intToRoman(num int) string {

	keys := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	vals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		val := vals[i]
		for num >= key {
			num -= key
			res += val
		}
	}

	return res
}

func Test12(t *testing.T) {
	//MMMDCCXLIX
	num := 1994
	res := intToRoman(num)
	fmt.Println(res)
}
