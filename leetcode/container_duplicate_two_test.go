package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsNearByDuplicates1(t *testing.T) {
	nums, k := []int{1, 2, 3, 1}, 3

	expected := true

	res := containsNearbyDuplicate(nums, k)

	assert.Equal(t, expected, res)
}

func TestContainsNearByDuplicates2(t *testing.T) {
	nums, k := []int{1, 0, 1, 1}, 1

	expected := true

	res := containsNearbyDuplicate(nums, k)

	assert.Equal(t, expected, res)
}

func TestContainsNearByDuplicates3(t *testing.T) {
	nums, k := []int{1, 2, 3, 1, 2, 3}, 2

	expected := false

	res := containsNearbyDuplicate(nums, k)

	assert.Equal(t, expected, res)
}
