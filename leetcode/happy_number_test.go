package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHappyNumber1(t *testing.T) {

	input := 19

	expected := true

	res := isHappy(input)

	assert.Equal(t, expected, res)
}

func TestHappyNumber2(t *testing.T) {

	input := 2

	expected := false

	res := isHappy(input)

	assert.Equal(t, expected, res)
}
