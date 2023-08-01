package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {

	nums_size := len(nums)

	if nums_size < 3 {
		return [][]int{}
	}

	sort.Ints(nums)

	cnt := 0

	res := [][]int{}

	for cnt < nums_size-2 {
		if cnt > 0 && nums[cnt] == nums[cnt-1] {
			cnt += 1
			continue
		}

		a := cnt + 1
		b := nums_size - 1
		target := 0 - nums[cnt]

		for a < b {
			if a > cnt+1 && nums[a] == nums[a-1] {
				a += 1
				continue
			}

			if b < nums_size-1 && nums[b] == nums[b+1] {
				b -= 1
				continue
			}

			tmp := nums[a] + nums[b]

			if tmp == target {
				res = append(res, []int{nums[cnt], nums[a], nums[b]})
			}

			if tmp > target {
				b -= 1
			} else {
				a += 1
			}
		}

		cnt += 1

	}

	return res
}
