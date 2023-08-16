package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLast1(t *testing.T) {

	ns := []int{1, 2, 3, 4, 5}

	assert.Equal(t, 5, getLast(ns))
}

func TestGetLast2(t *testing.T) {

	ss := []string{"first", "second", "last"}

	assert.Equal(t, "last", getLast(ss))
}

func TestGetLast3(t *testing.T) {

	ss := []string{}

	assert.Equal(t, "", getLast(ss))
}
