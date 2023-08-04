package leetcode

func candy(ratings []int) int {

	if len(ratings) == 0 {
		return 0
	}

	res := make([]int, len(ratings))
	res[0] = 1

	for cnt := 1; cnt < len(ratings); cnt++ {
		if ratings[cnt] > ratings[cnt-1] {
			res[cnt] = res[cnt-1] + 1
		} else {
			res[cnt] = 1
		}
	}

	numCandies := res[len(ratings)-1]

	for cnt := len(ratings) - 2; cnt >= 0; cnt-- {

		tmpCandy := 1

		if ratings[cnt] > ratings[cnt+1] {
			tmpCandy = res[cnt+1] + 1
		}

		numCandies += max(tmpCandy, res[cnt])
		res[cnt] = tmpCandy
	}

	return numCandies
}

func max(a int, b int) int {

	if a > b {
		return a
	}

	return b

}
