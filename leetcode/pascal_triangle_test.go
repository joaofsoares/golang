package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPascalTriangle1(t *testing.T) {
	expected := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}

	res := generate(5)

	assert.Equal(t, expected, res)
}

func TestPascalTriangle2(t *testing.T) {
	expected := [][]int{{1}}

	res := generate(1)

	assert.Equal(t, expected, res)
}
