package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameOfLife1(t *testing.T) {
	input := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}

	expected := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}

	gameOfLife(input)

	assert.Equal(t, expected, input)
}

func TestGameOfLife2(t *testing.T) {
	input := [][]int{{1, 1}, {1, 0}}

	expected := [][]int{{1, 1}, {1, 1}}

	gameOfLife(input)

	assert.Equal(t, expected, input)
}
