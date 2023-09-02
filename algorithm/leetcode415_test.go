package algorithm

import (
	"fmt"
	"strconv"
	"testing"
)

func addStrings(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return ""
	}
	n1 := len(num1) - 1
	n2 := len(num2) - 1

	m1 := 0
	res := ""
	for n1 >= 0 && n2 >= 0 {
		s1 := num1[n1] - '0'
		s2 := num2[n2] - '0'
		s3 := s1 + s2 + uint8(m1)
		m1 = int(s3 / 10)
		s := strconv.Itoa(int(s3 % 10))
		fmt.Println(s)
		res = s + res
		n1--
		n2--
	}
	for n1 >= 0 {
		s1 := num1[n1] - '0'
		s3 := s1 + uint8(m1)
		m1 = int(s3 / 10)
		res = strconv.Itoa(int(s3%10)) + res
		n1--
	}
	for n2 >= 0 {
		s2 := num2[n2] - '0'
		s3 := s2 + uint8(m1)
		m1 = int(s3 / 10)
		res = strconv.Itoa(int(s3%10)) + res
		n2--
	}
	if m1 == 1 {
		res = strconv.Itoa(m1) + res
	}
	return res
}

func Test415(t *testing.T) {
	num := addStrings("99", "99")
	fmt.Println(num)
}
