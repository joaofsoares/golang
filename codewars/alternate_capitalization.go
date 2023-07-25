package codewars

import "unicode"

func Capitalize(st string) []string {

	even, odd := []rune{}, []rune{}

	for i, c := range st {

		if i%2 == 0 {
			even = append(even, unicode.ToUpper(c))
			odd = append(odd, unicode.ToLower(c))
		} else {
			even = append(even, unicode.ToLower(c))
			odd = append(odd, unicode.ToUpper(c))
		}

	}

	return []string{string(even), string(odd)}
}
