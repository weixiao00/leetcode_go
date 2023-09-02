package algorithm

import (
	"fmt"
	"strings"
	"testing"
)

func canChange(start string, target string) bool {
	if strings.Replace(start, "_", "", -1) != strings.Replace(target, "_", "", -1) {
		return false
	}

	j := 0
	for i := 0; i < len(start); i++ {
		if start[i] == '_' {
			continue
		}
		for ; j < len(target); j++ {
			if target[j] != '_' {
				break
			}
		}
		if start[i] == 'L' && i < j {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}
		j++
	}

	return true
}

func Test2337(t *testing.T) {
	start := "_L__R__R_"
	target := "L______RR"
	change := canChange(start, target)
	fmt.Println(change)
}
