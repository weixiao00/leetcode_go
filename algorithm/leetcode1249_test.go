package algorithm

import (
	"fmt"
	"testing"
)

func minRemoveToMakeValid(s string) string {

	if len(s) == 0 {
		return s
	}

	arr := make([]int, 0)

	res := make([]bool, len(s))

	for i, ss := range s {
		if ss != '(' && ss != ')' {
			res[i] = true
			continue
		}
		if ss == '(' {
			arr = append(arr, i)
			continue
		}
		if len(arr) != 0 {
			res[i] = true
			res[arr[len(arr)-1]] = true
			arr = arr[0 : len(arr)-1]
		}
	}

	str := ""
	for i := range res {
		if res[i] {
			str += string(s[i])
		}
	}

	return str
}

func Test1249(t *testing.T) {
	//"lee(t(co)de)" , "lee(t(c)ode)"
	//leetc()o()de
	s := "a)b(c)d"
	res := minRemoveToMakeValid(s)
	fmt.Println(res)
}
