package exercism

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	strSize := len(id)
	if strSize <= 1 {
		return false
	}

	snd := false
	sum := 0

	for i := strSize - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}

		if snd {
			tmp := n * 2
			if tmp > 9 {
				tmp -= 9
			}

			sum += tmp
			snd = !snd
		} else {
			sum += n
			snd = !snd
		}
	}

	return sum%10 == 0
}
