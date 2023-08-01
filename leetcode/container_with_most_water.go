package leetcode

import "learn/helper"

func maxArea(height []int) int {

	left, right := 0, len(height)-1

	var res int

	for left < right {
		res = helper.Max(res, helper.Min(height[left], height[right])*(right-left))

		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}

	}

	return res
}
