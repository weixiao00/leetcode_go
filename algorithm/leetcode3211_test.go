package algorithm

import (
	"fmt"
	"testing"
)

// dfs
// 子字符串不包含两个相邻的0
func validStrings(n int) []string {

	res := make([]string, 0)

	if n == 0 {
		return res
	}

	var dfs func(str string)
	dfs = func(str string) {
		if len(str) == n {
			res = append(res, str)
			return
		}
		if len(str) == 0 || str[len(str)-1] == '1' {
			str += "0"
			dfs(str)
			str = str[0 : len(str)-1]
		}

		str += "1"
		dfs(str)
		str = str[0 : len(str)-1]
	}

	dfs("")
	return res
}

func Test3211(t *testing.T) {
	res := validStrings(1)
	fmt.Println(res)
}
