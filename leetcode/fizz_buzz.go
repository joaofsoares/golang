package leetcode

import "fmt"

func fizzBuzz(n int) []string {

	res := []string{}

	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			res = append(res, "FizzBuzz")
		case i%5 == 0:
			res = append(res, "Buzz")
		case i%3 == 0:
			res = append(res, "Fizz")
		default:
			res = append(res, fmt.Sprint(i))
		}
	}

	return res
}
