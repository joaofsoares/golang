package leetcode

func permute(nums []int) [][]int {

	res := [][]int{}

	permuteHelper(nums, len(nums), &res)

	return res
}

func permuteHelper(nums []int, l int, res *[][]int) {

	if l == 1 {
		*res = append(*res, append([]int{}, nums...))
		return
	}

	for i := 0; i < l; i++ {
		permuteHelper(nums, l-1, res)

		if l%2 == 0 {
			nums[i], nums[l-1] = nums[l-1], nums[i]
		} else {
			nums[0], nums[l-1] = nums[l-1], nums[0]
		}
	}

}
