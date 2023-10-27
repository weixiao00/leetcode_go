package algorithm

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

//给你一个正整数 num ，请你将它分割成两个非负整数 num1 和 num2 ，满足：
//
//num1 和 num2 直接连起来，得到 num 各数位的一个排列。
//<ul>
//	<li>换句话说，<code>num1</code> 和 <code>num2</code> 中所有数字出现的次数之和等于 <code>num</code> 中所有数字出现的次数。</li>
//</ul>
//</li>
//<li><code>num1</code> 和 <code>num2</code> 可以包含前导 0 。</li>
//请你返回 num1 和 num2 可以得到的和的 最小 值。
//
//注意：
//
//num 保证没有前导 0 。
//num1 和 num2 中数位顺序可以与 num 中数位顺序不同。
func splitNum(num int) int {
	numStr := strconv.Itoa(num)
	strArr := strings.Split(numStr, "")
	s := func(strArr []string) {
		sort.Slice(strArr, func(i, j int) bool {
			return strArr[i] < strArr[j]
		})
	}
	s(strArr)

	num1 := 0
	num2 := 0
	for i := 0; i < len(strArr); i++ {
		numInt, _ := strconv.Atoi(strArr[i])
		if i%2 == 0 {
			num1 = num1*10 + numInt
		} else {
			num2 = num2*10 + numInt
		}
	}
	return num1 + num2
}

// 官方解法
func splitNum1(num int) int {
	stnum := []byte(strconv.Itoa(num))
	sort.Slice(stnum, func(i, j int) bool {
		return stnum[i] < stnum[j]
	})
	num1, num2 := 0, 0
	for i := 0; i < len(stnum); i++ {
		if i%2 == 0 {
			num1 = num1*10 + int(stnum[i]-'0')
		} else {
			num2 = num2*10 + int(stnum[i]-'0')
		}
	}
	return num1 + num2
}

func Test2578(t *testing.T) {
	//num := 4325
	num1 := 11
	i := splitNum(num1)
	fmt.Println(i)
}
