package exercism

import "strings"

func Remove(str string) string {
	strSize := len(str)

	if strSize <= 1 {
		return str
	}

	if strSize == 2 {
		if str[0] == str[1] {
			return ""
		}
		return str
	}

	startAcc := removeHelper(str)
	nextAcc := removeHelper(startAcc)
	for isEqual := false; isEqual; isEqual = startAcc == nextAcc {
		startAcc = nextAcc
		nextAcc = removeHelper(startAcc)
	}

	return nextAcc
}

func removeHelper(str string) string {
	var acc strings.Builder
	strSize := len(str)
	i := 0

	for i < strSize {
		if i+1 < strSize && str[i] == str[i+1] {
			idx := i
			for isEqual := true; isEqual; isEqual = str[i] == str[idx] { // while in go lang
				idx++
			}
			i = idx
		} else {
			acc.WriteString(string(str[i]))
			i++
		}
	}

	return acc.String()
}
