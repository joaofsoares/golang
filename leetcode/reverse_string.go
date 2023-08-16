package leetcode

func reverseString(s []byte) {

	start, end := 0, len(s)-1

	for start < end {
		tmp := s[start]
		s[start] = s[end]
		s[end] = tmp

		start++
		end--
	}

}
