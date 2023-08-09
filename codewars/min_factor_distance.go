package codewars

import (
	"math"
	"sort"
)

func MinDistance(n int) int {

	nums := []int{}

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if n/i != i {
				nums = append(nums, i, n/i)
			}
		}
	}

	sort.Ints(nums)

	cnt, res := 0, 0

	for cnt < len(nums)-1 {

		if cnt == 0 {
			res = nums[cnt+1] - nums[cnt]
		}

		res = min(res, nums[cnt+1]-nums[cnt])

		cnt++
	}

	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
