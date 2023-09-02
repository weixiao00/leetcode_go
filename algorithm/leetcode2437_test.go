package algorithm

import (
	"fmt"
	"testing"
)

func Test2437(t *testing.T) {
	time := "0?:0?"
	countTime(time)
}

func countTime(time string) int {
	time1 := []byte(time)
	res := 0
	if len(time1) == 0 {
		return res
	}
	dfs(time1, 0, &res)
	fmt.Println(res)
	return res
}

func dfs(time []byte, pos int, res *int) {
	if pos == len(time) {
		if checkTime(time) {
			*res++
		}
		return
	}

	if time[pos] == '?' {
		for i := 0; i < 10; i++ {
			time[pos] = byte(i) + '0'
			dfs(time, pos+1, res)
			time[pos] = '?'
		}
	} else {
		dfs(time, pos+1, res)
	}
}

func checkTime(time []byte) bool {
	hh := (time[0]-'0')*10 + (time[1] - '0')
	mm := (time[3]-'0')*10 + (time[4] - '0')
	return (hh < 24) && (mm < 60)
}
