package codewars

import (
	"math"
	"strconv"
)

func DigPow(n int, p int) int {

	nums := convertIntToArray(n)

	sum := 0

	for cnt := 0; cnt < len(strconv.Itoa(n)); cnt++ {

		sum += int(math.Pow(float64(nums[cnt]), float64(p)))

		p++
	}

	res := sum / n

	if n*res == sum {
		return res
	}

	return -1
}

func convertIntToArray(num int) []int {

	nums := []int{}

	str := strconv.Itoa(num)

	for _, s := range str {
		n, err := strconv.Atoi(string(s))

		if err != nil {
			return []int{}
		}

		nums = append(nums, n)
	}

	return nums
}
