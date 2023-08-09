package codewars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicTacToeChecker1(t *testing.T) {
	input := [3][3]int{
		{0, 0, 1},
		{0, 1, 2},
		{2, 1, 0},
	}

	expected := -1

	res := IsSolved(input)

	assert.Equal(t, expected, res)
}

func TestTicTacToeChecker2(t *testing.T) {
	input := [3][3]int{
		{1, 1, 1},
		{0, 2, 2},
		{0, 0, 0},
	}

	expected := 1

	res := IsSolved(input)

	assert.Equal(t, expected, res)
}

func TestTicTacToeChecker3(t *testing.T) {
	input := [3][3]int{
		{2, 1, 2},
		{2, 1, 1},
		{1, 1, 2},
	}

	expected := 1

	res := IsSolved(input)

	assert.Equal(t, expected, res)
}

func TestTicTacToeChecker4(t *testing.T) {
	input := [3][3]int{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 1},
	}

	expected := 0

	res := IsSolved(input)

	assert.Equal(t, expected, res)
}

func TestTicTacToeChecker5(t *testing.T) {
	input := [3][3]int{
		{1, 1, 2},
		{0, 0, 0},
		{1, 2, 2},
	}

	expected := -1

	res := IsSolved(input)

	assert.Equal(t, expected, res)
}
