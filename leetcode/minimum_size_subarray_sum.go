package leetcode

import "learn/helper"

func minSubArrayLen(target int, nums []int) int {

	left, right := 0, -1
	tmpSum := 0
	minLen := len(nums) + 1

	for left < len(nums) {

		if right+1 < len(nums) && tmpSum < target {
			right += 1
			tmpSum += nums[right]
		} else {
			tmpSum -= nums[left]
			left += 1
		}

		if tmpSum >= target {
			minLen = helper.Min(minLen, right-left+1)
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}
