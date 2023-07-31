package leetcode

import (
	"strings"
)

func ZigZagConversion(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	m := make([][]string, numRows)

	rows := 0
	rev := false

	for _, v := range s {

		m[rows] = append(m[rows], string(v))

		if rows == numRows-1 && !rev {
			rev = true
			rows -= 1
		} else if rows == 0 && rev {
			rev = false
			rows += 1
		} else if rev {
			rows -= 1
		} else {
			rows += 1
		}

	}

	var res string

	for _, v := range m {
		res = res + strings.Join(v, "")
	}

	return res
}
