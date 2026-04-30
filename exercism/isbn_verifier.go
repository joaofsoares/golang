package exercism

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbnSize := len(isbn) - 1

	if isbnSize != 9 {
		return false
	}

	multiplier := 1
	var sum int

	for isbnSize >= 0 {
		value := rune(isbn[isbnSize])

		if isbnSize == 9 && value == 'X' {
			sum += 10
		} else if unicode.IsDigit(value) {
			sum += int(value-'0') * multiplier
		} else {
			return false
		}

		multiplier++
		isbnSize--
	}

	return sum%11 == 0
}
