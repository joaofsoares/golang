package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutate1(t *testing.T) {
	input := []int{1, 2, 3}

	expected := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}

	res := permute(input)

	assert.Equal(t, len(expected), len(res))
}

func TestPermutate2(t *testing.T) {
	input := []int{0, 1}

	expected := [][]int{{0, 1}, {1, 0}}

	res := permute(input)

	assert.Equal(t, expected, res)
}

func TestPermutate3(t *testing.T) {
	input := []int{1}

	expected := [][]int{{1}}

	res := permute(input)

	assert.Equal(t, expected, res)
}
