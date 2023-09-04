package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDigits1(t *testing.T) {
	assert.Equal(t, 2, addDigits(38))
}

func TestAddDigits2(t *testing.T) {
	assert.Equal(t, 0, addDigits(0))
}
