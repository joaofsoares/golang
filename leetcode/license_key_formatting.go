package leetcode

import (
	"strings"
)

func licenseKeyFormatting(s string, k int) string {

	s = strings.ReplaceAll(strings.ToUpper(s), "-", "")

	tmp := ""
	res := ""

	for i := len(s) - 1; i >= 0; i-- {
		tmp = string(s[i]) + tmp

		if len(tmp) == k {
			if res == "" {
				res = tmp
			} else {
				res = tmp + "-" + res
			}
			tmp = ""
		}
	}

	if tmp != "" && res != "" {
		return tmp + "-" + res
	} else if tmp != "" && res == "" {
		return tmp
	} else {
		return res
	}
}
