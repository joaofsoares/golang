package exercism

import "errors"

func Distance(a, b string) (int, error) {
	var cnt int = 0
	var size int = len(a)

	if size != len(b) {
		return cnt, errors.New("disallow different strings")
	}

	for i := range size {
		if (a[i] != b[i]) {
			cnt++
		}
	}

	return cnt, nil
}
