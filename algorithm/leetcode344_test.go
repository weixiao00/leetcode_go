package algorithm

import (
	"fmt"
	"testing"
)

func reverseString(s []byte) {
	fmt.Println(string(s))
	begin := 0
	len := len(s)
	for begin < len {
		exchange(&s[begin], &s[len-1])
		begin++
		len--
	}
	fmt.Println(string(s))
}

func exchange(a, b *byte) {
	temp := *a
	*a = *b
	*b = temp
}

func Test344(t *testing.T) {
	s := "hello"
	reverseString([]byte(s))
}
