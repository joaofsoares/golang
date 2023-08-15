package leetcode

func generate(numRows int) [][]int {

	res := make([][]int, numRows)

	for i := 1; i <= numRows; i++ {
		col := 1
		tmpLine := []int{}

		for j := 1; j <= i; j++ {
			tmpLine = append(tmpLine, col)
			col = col * (i - j) / j
		}

		res[i-1] = tmpLine
	}

	return res
}
