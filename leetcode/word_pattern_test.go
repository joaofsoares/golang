package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPatternWord1(t *testing.T) {
	pattern, input := "abba", "dog cat cat dog"

	expected := true

	res := wordPattern(pattern, input)

	assert.Equal(t, expected, res)
}

func TestPatternWord2(t *testing.T) {
	pattern, input := "abba", "dog cat cat fish"

	expected := false

	res := wordPattern(pattern, input)

	assert.Equal(t, expected, res)
}

func TestPatternWord3(t *testing.T) {
	pattern, input := "aaaa", "dog cat cat dog"

	expected := false

	res := wordPattern(pattern, input)

	assert.Equal(t, expected, res)
}

func TestPatternWord4(t *testing.T) {
	pattern, input := "abba", "dog dog dog dog"

	expected := false

	res := wordPattern(pattern, input)

	assert.Equal(t, expected, res)
}

func TestPatternWord5(t *testing.T) {
	pattern, input := "aaa", "aa aa aa aa"

	expected := false

	res := wordPattern(pattern, input)

	assert.Equal(t, expected, res)
}
