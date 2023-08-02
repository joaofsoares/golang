package leetcode

import "strings"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[byte]rune)
	cs := ""

	cnt := 0

	for cnt < len(s) {

		sVal, sExist := m[s[cnt]]
		tVal := rune(t[cnt])

		if sExist {
			if sVal != tVal {
				return false
			}
		} else {

			if strings.ContainsRune(cs, tVal) {
				return false
			}

			m[s[cnt]] = tVal
			cs += string(tVal)

		}

		cnt += 1
	}

	return true
}
