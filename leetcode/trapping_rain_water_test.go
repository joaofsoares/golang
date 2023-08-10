package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap1(t *testing.T) {

	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	expected := 6

	res := trap(input)

	assert.Equal(t, expected, res)
}

func TestTrap2(t *testing.T) {

	input := []int{4, 2, 0, 3, 2, 5}

	expected := 9

	res := trap(input)

	assert.Equal(t, expected, res)
}
