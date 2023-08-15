package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateList1(t *testing.T) {

	input, n := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2

	expected := &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}

	res := rotateRight(input, n)

	assert.Equal(t, expected, res)
}

func TestRotateList2(t *testing.T) {

	input, n := &ListNode{0, &ListNode{1, &ListNode{2, nil}}}, 4

	expected := &ListNode{2, &ListNode{0, &ListNode{1, nil}}}

	res := rotateRight(input, n)

	assert.Equal(t, expected, res)
}

func TestRotateList3(t *testing.T) {

	input, n := &ListNode{1, &ListNode{2, nil}}, 1

	expected := &ListNode{2, &ListNode{1, nil}}

	res := rotateRight(input, n)

	assert.Equal(t, expected, res)
}

func TestRotateList4(t *testing.T) {

	input, n := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2000000000

	expected := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	res := rotateRight(input, n)

	assert.Equal(t, expected, res)
}

func TestRotateList5(t *testing.T) {

	input, n := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}, 2000000000

	expected := &ListNode{2, &ListNode{3, &ListNode{1, nil}}}

	res := rotateRight(input, n)

	assert.Equal(t, expected, res)
}
