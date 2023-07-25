package codewars

func Gimme(array [3]int) int {

	for i, v := range array {
		if i == 0 {

			if (v < array[i+1] && v > array[i+2]) || (v > array[i+1] && v < array[i+2]) {
				return i
			}

		} else if i == 2 {

			if (v < array[i-1] && v > array[i-2]) || (v > array[i-1] && v < array[i-2]) {
				return i
			}

		} else {

			if (v < array[i-1] && v > array[i+1]) || (v > array[i-1] && v < array[i+1]) {
				return i
			}

		}
	}

	return 0
}
