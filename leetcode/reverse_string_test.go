package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString1(t *testing.T) {
	s := []byte("hello")

	expected := []byte("olleh")

	reverseString([]byte(s))

	assert.Equal(t, expected, s)
}

func TestReverseString2(t *testing.T) {
	s := []byte("Hannah")

	expected := []byte("hannaH")

	reverseString([]byte(s))

	assert.Equal(t, expected, s)
}
