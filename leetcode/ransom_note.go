package leetcode

import (
	"sort"
)

func sortString(s *string) {
	chars := []rune(*s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j] //sort the string rune
	})
	*s = string(chars)
}

func canConstruct(ransomNote string, magazine string) bool {

	sortString(&ransomNote)
	sortString(&magazine)

	cntMagazine, cntRansomNote := 0, 0

	for cntMagazine < len(magazine) {

		if cntRansomNote < len(ransomNote) && ransomNote[cntRansomNote] == magazine[cntMagazine] {
			cntMagazine += 1
			cntRansomNote += 1
		} else {
			cntMagazine += 1
		}

		if cntRansomNote == (len(ransomNote)) {
			return true
		}

	}

	return cntRansomNote == (len(ransomNote))
}
