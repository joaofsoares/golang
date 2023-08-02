package leetcode

func isHappy(n int) bool {

	slow, fast := n, n

	for {
		slow = squareSum(slow)
		fast = squareSum(squareSum(fast))

		if slow == fast {
			break
		}
	}

	return slow == 1
}

func squareSum(num int) int {
	tmpSum := 0

	for num > 0 {
		tmpSum += (num % 10) * (num % 10)
		num = num / 10
	}

	return tmpSum
}
