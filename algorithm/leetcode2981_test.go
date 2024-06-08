package algorithm

import (
	"fmt"
	"testing"
)

// 只包含小写字母
func maximumLength(s string) int {
	if len(s) == 0 {
		return 0
	}
	chas := make([][]int, 26)
	for i := range chas {
		chas[i] = make([]int, 0)
	}

	cnt := 0
	for i, _ := range s {
		cnt++
		if i == len(s)-1 || s[i] != s[i+1] {
			chaArr := chas[s[i]-'a']
			chaArr = append(chaArr, cnt)
			chas[s[i]-'a'] = chaArr
			cnt = 0

			// 排序
			for i := len(chaArr) - 1; i > 0; i-- {
				if chaArr[i] > chaArr[i-1] {
					swap(chaArr, i, i-1)
				}
			}
			// 只要三个就行
			if len(chaArr) > 3 {
				chaArr = chaArr[0:3]
			}
		}
	}

	res := -1
	for _, chaArr := range chas {
		if len(chaArr) > 0 && chaArr[0] > 2 {
			res = getMax7(res, chaArr[0]-2)
		}
		if len(chaArr) > 1 && chaArr[0] > 1 {
			res = getMax7(res, getMin1(chaArr[0]-1, chaArr[1]))
		}
		if len(chaArr) > 2 {
			res = getMax7(res, chaArr[2])
		}
	}

	return res
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func getMax7(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin1(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test2981(t *testing.T) {
	s := "abcdef"
	length := maximumLength(s)
	fmt.Println(length)
}
