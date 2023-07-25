package helper

type Comp interface {
	int | string
}

func AreEquals[A Comp](arr1 []A, arr2 []A) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}

	return true
}
