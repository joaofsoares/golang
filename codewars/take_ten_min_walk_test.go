package codewars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeWalk1(t *testing.T) {

	input := []rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}

	expected := true

	res := IsValidWalk(input)

	assert.Equal(t, expected, res)
}

func TestTakeWalk2(t *testing.T) {

	input := []rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'}

	expected := false

	res := IsValidWalk(input)

	assert.Equal(t, expected, res)
}

func TestTakeWalk3(t *testing.T) {

	input := []rune{'w'}

	expected := false

	res := IsValidWalk(input)

	assert.Equal(t, expected, res)
}

func TestTakeWalk4(t *testing.T) {

	input := []rune{'n', 'n', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}

	expected := false

	res := IsValidWalk(input)

	assert.Equal(t, expected, res)
}

func TestTakeWalk5(t *testing.T) {

	input := []rune{'e', 'e', 'e', 'e', 'w', 'w', 's', 's', 's', 's'}

	expected := false

	res := IsValidWalk(input)

	assert.Equal(t, expected, res)
}
