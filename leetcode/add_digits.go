package leetcode

import "strconv"

func addDigits(num int) int {

	s := strconv.Itoa(num)

	if len(s) == 1 {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return i
	}

	sum := 0

	for _, v := range s {
		sum += int(v - '0')
	}

	return addDigits(sum)
}
