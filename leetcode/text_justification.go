package leetcode

import (
	"math"
	"strings"
)

func formatValidStr(validStr *[]string, maxWidth *int) string {

	concatStr := strings.Join(*validStr, "")

	diff := *maxWidth - len(concatStr)

	if len(*validStr) == 1 {
		return strings.Join(*validStr, " ") + strings.Repeat(" ", diff)
	} else {
		numSep, r := divMod(diff, (len(*validStr) - 1))
		return createStr(validStr, numSep, r)
	}
}

func createStr(validStr *[]string, numSep int, rest int) string {

	var ret string

	for i, v := range *validStr {
		if rest > 0 {
			ret = ret + v + strings.Repeat(" ", numSep+1)
			rest -= 1
		} else if i == len(*validStr)-1 {
			ret = ret + v
		} else {
			ret = ret + v + strings.Repeat(" ", numSep)
		}
	}

	return ret

}

func divMod(num int, div int) (int, int) {
	d := num / div
	m := int(math.Mod(float64(num), float64(div)))
	return d, m

}

func formatLastWord(last *[]string, maxWidth *int) string {
	formatStr := strings.Join(*last, " ")
	diff := *maxWidth - len(formatStr)
	return formatStr + strings.Repeat(" ", diff)
}

func fullJustify(words []string, maxWidth int) []string {

	cnt := 0

	var res []string
	var tmpStr []string

	for cnt < len(words) {

		str := words[cnt]
		tmpStr = append(tmpStr, str)
		ss := strings.Join(tmpStr, " ")

		var validStr string

		if len(ss) > maxWidth {
			validStrs := tmpStr[:len(tmpStr)-1]

			validStr = formatValidStr(&validStrs, &maxWidth)
			res = append(res, validStr)

			tmpStr = append([]string{}, tmpStr[len(tmpStr)-1])
		}

		cnt += 1
	}

	if tmpStr != nil {
		last := formatLastWord(&tmpStr, &maxWidth)
		res = append(res, last)
	}

	return res
}
