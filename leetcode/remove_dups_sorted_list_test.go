package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDups1(t *testing.T) {
	input := &ListNode{1, &ListNode{1, &ListNode{2, nil}}}

	expected := &ListNode{1, &ListNode{2, nil}}

	res := deleteDuplicates(input)

	assert.Equal(t, expected, res)
}

func TestDeleteDups2(t *testing.T) {
	input := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}

	expected := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}

	res := deleteDuplicates(input)

	assert.Equal(t, expected, res)
}
