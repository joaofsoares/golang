package exercism

import "errors"

func CollatzConjecture(n int) (int, error) {
	var cnt int = 0

	if n <= 0 {
		return cnt, errors.New("zero or negative numbers are not valid")
	}

	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = (n * 3) + 1
		}

		cnt++
	}

	return cnt, nil
}
