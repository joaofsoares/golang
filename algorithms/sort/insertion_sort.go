package sort

func InsertionSort(arr *[]int, size int) {
	for i := 1; i < size; i++ {
		key := (*arr)[i]
		j := i - 1

		for j >= 0 && (*arr)[j] > key {
			(*arr)[j+1] = (*arr)[j]
			j = j - 1
		}

		(*arr)[j+1] = key
	}
}
