package codewars

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinDistance1(t *testing.T) {
	input := 8
	expected := 1

	res := MinDistance(input)

	assert.Equal(t, expected, res)
}

func TestMinDistance2(t *testing.T) {
	input := 13013
	expected := 2

	res := MinDistance(input)

	assert.Equal(t, expected, res)
}

func TestMinDistance3(t *testing.T) {
	input := 408593
	expected := 352

	res := MinDistance(input)

	assert.Equal(t, expected, res)
}

func TestMinDistance4(t *testing.T) {
	input := 1
	expected := 0

	res := MinDistance(input)

	assert.Equal(t, expected, res)
}
