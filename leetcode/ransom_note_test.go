package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanConstruct1(t *testing.T) {
	input1, input2 := "a", "b"

	expected := false

	res := canConstruct(input1, input2)

	assert.Equal(t, expected, res)
}

func TestCanConstruct2(t *testing.T) {
	input1, input2 := "aa", "ab"

	expected := false

	res := canConstruct(input1, input2)

	assert.Equal(t, expected, res)
}

func TestCanConstruct3(t *testing.T) {
	input1, input2 := "aa", "aab"

	expected := true

	res := canConstruct(input1, input2)

	assert.Equal(t, expected, res)
}
