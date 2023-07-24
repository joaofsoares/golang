package codewars

import "strings"

func DuplicateCount(s string) int {

	dups := make(map[rune]int)

	for _, v := range strings.ToLower(s) {
		tmpValue, exists := dups[v]

		if exists {
			dups[v] = tmpValue + 1
		} else {
			dups[v] = 1
		}
	}

	cnt := 0

	for _, v := range dups {
		if v > 1 {
			cnt++
		}
	}

	return cnt
}
