package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddStrings1(t *testing.T) {
	s1, s2 := "11", "123"

	expected := "134"

	res := addStrings(s1, s2)

	assert.Equal(t, expected, res)
}

func TestAddStrings2(t *testing.T) {
	s1, s2 := "456", "77"

	expected := "533"

	res := addStrings(s1, s2)

	assert.Equal(t, expected, res)
}

func TestAddStrings3(t *testing.T) {
	s1, s2 := "0", "0"

	expected := "0"

	res := addStrings(s1, s2)

	assert.Equal(t, expected, res)
}

func TestAddStrings4(t *testing.T) {
	s1, s2 := "9", "99"

	expected := "108"

	res := addStrings(s1, s2)

	assert.Equal(t, expected, res)
}
