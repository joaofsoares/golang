package codewars

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeastCommonMultiple1(t *testing.T) {
	expected := big.NewInt(10)

	res := LCM(2, 5)

	assert.Equal(t, expected, res)
}

func TestLeastCommonMultiple2(t *testing.T) {
	expected := big.NewInt(12)

	res := LCM(2, 3, 4)

	assert.Equal(t, expected, res)
}

func TestLeastCommonMultiple3(t *testing.T) {
	expected := big.NewInt(9)

	res := LCM(9)

	assert.Equal(t, expected, res)
}

func TestLeastCommonMultiple4(t *testing.T) {
	expected := big.NewInt(0)

	res := LCM(0)

	assert.Equal(t, expected, res)
}

func TestLeastCommonMultiple5(t *testing.T) {
	expected := big.NewInt(0)

	res := LCM(0, 1)

	assert.Equal(t, expected, res)
}

func TestLeastCommonMultiple6(t *testing.T) {
	expected := big.NewInt(0)

	res := LCM(0, 1, 0)

	assert.Equal(t, expected, res)
}

func TestLeastCommonMultiple7(t *testing.T) {
	expected := big.NewInt(1)

	res := LCM(1, 1, 1)

	assert.Equal(t, expected, res)
}
