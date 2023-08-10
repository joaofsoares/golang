package leetcode

func trap(height []int) int {

	volume := 0

	left, right := make([]int, len(height)), make([]int, len(height))

	left[0] = height[0]
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}

	right[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	for i := 1; i < len(height)-1; i++ {
		tmp := min(left[i-1], right[i+1])
		if tmp > height[i] {
			volume += tmp - height[i]
		}
	}

	return volume
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
