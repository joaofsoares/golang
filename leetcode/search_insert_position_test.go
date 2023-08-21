package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert1(t *testing.T) {
	input, target := []int{1, 3, 5, 6}, 5

	res := searchInsert(input, target)

	assert.Equal(t, 2, res)
}

func TestSearchInsert2(t *testing.T) {
	input, target := []int{1, 3, 5, 6}, 2

	res := searchInsert(input, target)

	assert.Equal(t, 1, res)
}

func TestSearchInsert3(t *testing.T) {
	input, target := []int{1, 3, 5, 6}, 7

	res := searchInsert(input, target)

	assert.Equal(t, 4, res)
}

func TestSearchInsert4(t *testing.T) {
	input, target := []int{1, 3, 5, 6}, 4

	res := searchInsert(input, target)

	assert.Equal(t, 2, res)
}

func TestSearchInsert5(t *testing.T) {
	input, target := []int{1, 3, 5, 6}, 6

	res := searchInsert(input, target)

	assert.Equal(t, 3, res)
}
