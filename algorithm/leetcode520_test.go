package algorithm

import "unicode"

func detectCapitalUse(word string) bool {

	if len(word) < 2 {
		return true
	}

	first := word[0]
	second := word[1]

	if unicode.IsLower(rune(first)) && unicode.IsUpper(rune(second)) {
		return false
	}

	for i := 2; i < len(word); i++ {
		if unicode.IsLower(rune(second)) != unicode.IsLower(rune(word[i])) {
			return false
		}
	}
	return true
}
