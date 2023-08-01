package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea1(t *testing.T) {

	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	expected := 49

	result := maxArea(input)

	assert.Equal(t, expected, result)
}

func TestMaxArea2(t *testing.T) {

	input := []int{1, 1}

	expected := 1

	result := maxArea(input)

	assert.Equal(t, expected, result)
}
