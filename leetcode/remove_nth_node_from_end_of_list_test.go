package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveNth1(t *testing.T) {

	nodes, element := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2

	expected := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}}

	res := removeNthFromEnd(nodes, element)

	headExpected := expected
	headRes := res

	for headExpected != nil && headRes != nil {
		assert.Equal(t, headExpected.Val, headRes.Val)
		headExpected = headExpected.Next
		headRes = headRes.Next
	}

}

func TestRemoveNth2(t *testing.T) {

	nodes, element := &ListNode{1, nil}, 1

	expected := &ListNode{0, nil}

	res := removeNthFromEnd(nodes, element)

	assert.Equal(t, res, expected.Next)
}

func TestRemoveNth3(t *testing.T) {

	nodes, element := &ListNode{1, &ListNode{2, nil}}, 1

	expected := &ListNode{1, nil}

	res := removeNthFromEnd(nodes, element)

	headExpected := expected
	headRes := res

	for headExpected != nil && headRes != nil {
		assert.Equal(t, headExpected.Val, headRes.Val)
		headExpected = headExpected.Next
		headRes = headRes.Next
	}

}
