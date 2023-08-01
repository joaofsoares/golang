package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSubArray1(t *testing.T) {

	input, target := []int{2, 3, 1, 2, 4, 3}, 7

	expected := 2

	res := minSubArrayLen(target, input)

	assert.Equal(t, expected, res)
}

func TestMinSubArray2(t *testing.T) {

	input, target := []int{1, 4, 4}, 4

	expected := 1

	res := minSubArrayLen(target, input)

	assert.Equal(t, expected, res)
}

func TestMinSubArray3(t *testing.T) {

	input, target := []int{1, 1, 1, 1, 1, 1, 1, 1}, 11

	expected := 0

	res := minSubArrayLen(target, input)

	assert.Equal(t, expected, res)
}
