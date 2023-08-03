package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSummaryRanges1(t *testing.T) {
	input := []int{0, 1, 2, 4, 5, 7}
	expected := []string{"0->2", "4->5", "7"}
	res := summaryRanges(input)
	assert.Equal(t, expected, res)
}

func TestSummaryRanges2(t *testing.T) {
	input := []int{0, 2, 3, 4, 6, 8, 9}
	expected := []string{"0", "2->4", "6", "8->9"}
	res := summaryRanges(input)
	assert.Equal(t, expected, res)
}

func TestSummaryRanges3(t *testing.T) {
	input := []int{0, 2, 4, 6, 8}
	expected := []string{"0", "2", "4", "6", "8"}
	res := summaryRanges(input)
	assert.Equal(t, expected, res)
}
