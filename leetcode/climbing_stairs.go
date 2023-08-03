package leetcode

func climbStairs(n int) int {
	pre, curr, tmp := 1, 1, 0

	for i := 1; i < n; i++ {
		tmp = curr
		curr = curr + pre
		pre = tmp
	}

	return curr
}
