package codewars

import (
	"strings"
	"unicode"
)

func ToWeirdCase(s string) string {

	words := strings.Fields(s)

	for idx, word := range words {
		runeWord := []rune(word)
		for i, c := range runeWord {
			if i%2 == 0 {
				runeWord[i] = unicode.ToUpper(c)
			} else {
				runeWord[i] = unicode.ToLower(c)
			}
		}
		words[idx] = string(runeWord)
	}

	return strings.Join(words, " ")
}
