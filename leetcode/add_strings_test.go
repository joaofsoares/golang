package leetcode_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"learn/leetcode"
)

func TestAddStrings1(t *testing.T) {
	s1, s2 := "11", "123"

	expected := "134"

	res := leetcode.AddStrings(s1, s2)

	assert.Equal(t, expected, res)
}

func TestAddStrings2(t *testing.T) {
	s1, s2 := "456", "77"

	expected := "533"

	res := leetcode.AddStrings(s1, s2)

	assert.Equal(t, expected, res)
}

func TestAddStrings3(t *testing.T) {
	s1, s2 := "0", "0"

	expected := "0"

	res := leetcode.AddStrings(s1, s2)

	assert.Equal(t, expected, res)
}

func TestAddStrings4(t *testing.T) {
	s1, s2 := "9", "99"

	expected := "108"

	res := leetcode.AddStrings(s1, s2)

	assert.Equal(t, expected, res)
}

func TestAddStringsTable(t *testing.T) {
	data := []struct {
		s1       string
		s2       string
		expected string
	}{
		{s1: "11", s2: "123", expected: "134"},
		{s1: "456", s2: "77", expected: "533"},
		{s1: "0", s2: "0", expected: "0"},
		{s1: "9", s2: "99", expected: "108"},
	}

	for _, val := range data {
		res := leetcode.AddStrings(val.s1, val.s2)

		if res != val.expected {
			fmt.Errorf("s1 = %s + s2 = %s is not equal %s, instead received = %s",
				val.s1, val.s2, val.expected, res)
		}
	}

}
