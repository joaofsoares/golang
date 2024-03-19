package sum

func SumArray(arr *[]int, size int) (sum int) {
	for i := 0; i < size; i++ {
		sum = sum + (*arr)[i]
	}
	return sum
}
