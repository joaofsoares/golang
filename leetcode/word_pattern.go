package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {

	words := strings.Fields(s)

	if len(pattern) != len(words) {
		return false
	}

	m := map[rune]string{}

	for i, key := range pattern {
		w, exist := m[key]

		word := words[i]

		if exist && w != word {
			return false
		} else {
			for k, v := range m {
				if word == v && k != key {
					return false
				}
			}

			m[key] = word
		}
	}

	return true
}
