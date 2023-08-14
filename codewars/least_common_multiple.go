package codewars

import "math/big"

func LCM(nums ...int64) *big.Int {

	if len(nums) == 0 {
		return big.NewInt(1)
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			return big.NewInt(0)
		}
	}

	res := big.NewInt(nums[0])
	for i := 1; i < len(nums); i++ {
		gcd := big.NewInt(0).GCD(nil, nil, res, big.NewInt(nums[i]))
		res = res.Mul(res, big.NewInt(nums[i])).Div(res, gcd)
	}

	return res
}
