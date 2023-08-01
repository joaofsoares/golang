package leetcode

func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	} else if v2 > v1 {
		return v2
	} else {
		return v1
	}
}

func min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	} else if v2 < v1 {
		return v2
	} else {
		return v1
	}
}

func maxArea(height []int) int {

	left, right := 0, len(height)-1

	var res int

	for left < right {

		res = max(res, min(height[left], height[right])*(right-left))

		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}

	}

	return res
}
