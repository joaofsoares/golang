package leetcode

func mySqrt(x int) int {

	if x == 0 || x == 1 {
		return x
	}

	start, end, res := 1, x/2, 0

	for start <= end {
		mid := (start + end) / 2
		tmp := mid * mid

		if tmp == x {
			return mid
		}

		if tmp <= x {
			start = mid + 1
			res = mid
		} else {
			end = mid - 1
		}

	}

	return res

}
