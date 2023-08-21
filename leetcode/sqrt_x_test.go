package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrtX1(t *testing.T) {
	assert.Equal(t, 2, mySqrt(4))
}

func TestSqrtX2(t *testing.T) {
	assert.Equal(t, 2, mySqrt(8))
}
