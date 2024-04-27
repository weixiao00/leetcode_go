package algorithm

import (
	"fmt"
	"strconv"
	"testing"
)

func countSeniors(details []string) int {

	if len(details) == 0 {
		return 0
	}

	res := 0
	for _, d := range details {
		s := d[11:13]
		age, _ := strconv.Atoi(s)
		if age > 60 {
			res++
		}
	}

	return res
}

func Test2678(t *testing.T) {
	str := []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}
	seniors := countSeniors(str)
	fmt.Println(seniors)
}
