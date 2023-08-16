package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueChar1(t *testing.T) {
	s := "leetcode"

	assert.Equal(t, 0, firstUniqChar(s))
}

func TestUniqueChar2(t *testing.T) {
	s := "loveleetcode"

	assert.Equal(t, 2, firstUniqChar(s))
}

func TestUniqueChar3(t *testing.T) {
	s := "aabb"

	assert.Equal(t, -1, firstUniqChar(s))
}
