package leetcode

import "strconv"

func summaryRanges(nums []int) []string {
	res := []string{}

	cnt := 0

	for cnt < len(nums) {
		start, end := nums[cnt], 0

		for cnt+1 < len(nums) && nums[cnt+1] == nums[cnt]+1 {
			cnt += 1
		}
		end = nums[cnt]

		if start == end {
			res = append(res, strconv.Itoa(start))
		} else {
			res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(end))
		}

		cnt += 1
	}
	return res
}
