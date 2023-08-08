package codewars

import "unicode"

func ValidISBN10(isbn string) bool {

	if len(isbn) != 10 {
		return false
	}

	sum, cnt := 0, 0

	for cnt < len(isbn) {

		if cnt+1 == 10 && (isbn[cnt] == 'X' || isbn[cnt] == 'x') {
			sum += 10 * 10
			cnt++
			continue
		} else if !unicode.IsDigit(rune(isbn[cnt])) {
			return false
		}

		sum += (cnt + 1) * int(isbn[cnt]-'0')

		cnt++
	}

	if sum%11 == 0 {
		return true
	}

	return false
}
