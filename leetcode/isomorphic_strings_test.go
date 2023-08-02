package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsIsomorphic1(t *testing.T) {
	input1, input2 := "egg", "add"

	expected := true

	res := isIsomorphic(input1, input2)

	assert.Equal(t, expected, res)
}

func TestIsIsomorphic2(t *testing.T) {
	input1, input2 := "foo", "bar"

	expected := false

	res := isIsomorphic(input1, input2)

	assert.Equal(t, expected, res)
}

func TestIsIsomorphic3(t *testing.T) {
	input1, input2 := "paper", "title"

	expected := true

	res := isIsomorphic(input1, input2)

	assert.Equal(t, expected, res)
}

func TestIsIsomorphic4(t *testing.T) {
	input1, input2 := "badc", "baba"

	expected := false

	res := isIsomorphic(input1, input2)

	assert.Equal(t, expected, res)
}
