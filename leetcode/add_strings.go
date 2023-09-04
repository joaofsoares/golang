package leetcode

import "fmt"

func AddStrings(num1 string, num2 string) string {
	s1 := len(num1) - 1
	s2 := len(num2) - 1

	rem := 0
	res := ""

	for s1 > -1 && s2 > -1 {
		sum := rem + int(num1[s1]-'0') + int(num2[s2]-'0')
		checkValue(&res, &rem, sum)

		s1--
		s2--
	}

	for s1 > -1 {
		sum := rem + int(num1[s1]-'0')
		checkValue(&res, &rem, sum)

		s1--
	}

	for s2 > -1 {
		sum := rem + int(num2[s2]-'0')
		checkValue(&res, &rem, sum)

		s2--
	}

	if rem > 0 {
		res = "1" + res
	}

	return res
}

func checkValue(res *string, rem *int, sum int) {
	if sum > 9 {
		*rem = 1
		*res = fmt.Sprint(sum-10) + *res
	} else {
		*rem = 0
		*res = fmt.Sprint(sum) + *res
	}
}
