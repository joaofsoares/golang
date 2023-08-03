package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbingStairs1(t *testing.T) {
	input := 2
	expected := 2
	res := climbStairs(input)

	assert.Equal(t, expected, res)
}

func TestClimbingStairs2(t *testing.T) {
	input := 3
	expected := 3
	res := climbStairs(input)

	assert.Equal(t, expected, res)
}
