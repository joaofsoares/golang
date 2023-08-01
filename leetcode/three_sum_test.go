package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum1(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}

	res := threeSum(input)

	assert.Equal(t, expected, res)
}

func TestThreeSum2(t *testing.T) {
	input := []int{0, 1, 1}
	expected := [][]int{}

	res := threeSum(input)

	assert.Equal(t, expected, res)
}

func TestThreeSum3(t *testing.T) {
	input := []int{0, 0, 0}
	expected := [][]int{{0, 0, 0}}

	res := threeSum(input)

	assert.Equal(t, expected, res)
}
