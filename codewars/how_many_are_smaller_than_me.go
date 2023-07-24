package codewars

func Smaller(arr []int) []int {

	for i, v := range arr {
		cnt := 0
		for _, v2 := range arr[i:(len(arr))] {
			if v > v2 {
				cnt++
			}
		}
		arr[i] = cnt
	}

	return arr
}
