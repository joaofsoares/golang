package codewars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayingWithDigits1(t *testing.T) {

	num, p := 89, 1

	expected := 1

	res := DigPow(num, p)

	assert.Equal(t, expected, res)
}

func TestPlayingWithDigits2(t *testing.T) {

	num, p := 92, 1

	expected := -1

	res := DigPow(num, p)

	assert.Equal(t, expected, res)
}

func TestPlayingWithDigits3(t *testing.T) {

	num, p := 695, 2

	expected := 2

	res := DigPow(num, p)

	assert.Equal(t, expected, res)
}

func TestPlayingWithDigits4(t *testing.T) {

	num, p := 46288, 3

	expected := 51

	res := DigPow(num, p)

	assert.Equal(t, expected, res)
}
