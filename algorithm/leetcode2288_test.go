package algorithm

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func discountPrices(sentence string, discount int) string {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if strings.HasPrefix(word, "$") && isNumber(word[1:]) {
			number, _ := strconv.Atoi(word[1:])
			newNumber := float64(number) * (1 - float64(discount)/100)
			words[i] = "$" + fmt.Sprintf("%.2f", newNumber)
		}
	}

	return strings.Join(words, " ")
}

func isNumber(s string) bool {
	match, _ := regexp.MatchString("^[0-9]+$", s)
	return match
}

func Test2288(t *testing.T) {
	sentence := "there are $1 $2 and 5$ candies in the shop"
	discount := 50
	prices := discountPrices(sentence, discount)
	fmt.Println(prices)
}
