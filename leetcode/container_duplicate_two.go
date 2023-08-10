package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {

	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {

		v, exist := m[nums[i]]

		if exist && i-v <= k {
			return true
		}

		m[nums[i]] = i

	}

	return false
}
