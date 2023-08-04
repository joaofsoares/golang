package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCandy1(t *testing.T) {

	input := []int{1, 0, 2}

	expected := 5

	res := candy(input)

	assert.Equal(t, expected, res)
}

func TestCandy2(t *testing.T) {

	input := []int{1, 2, 2}

	expected := 4

	res := candy(input)

	assert.Equal(t, expected, res)
}

func TestCandy3(t *testing.T) {

	input := []int{29, 51, 87, 87, 72, 12}

	expected := 12

	res := candy(input)

	assert.Equal(t, expected, res)
}
