package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSameTree1(t *testing.T) {
	t1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	t2 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}

	assert.Equal(t, true, isSameTree(t1, t2))
}

func TestSameTree2(t *testing.T) {
	t1 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	t2 := &TreeNode{1, nil, &TreeNode{2, nil, nil}}

	assert.Equal(t, false, isSameTree(t1, t2))
}

func TestSameTree3(t *testing.T) {
	t1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil}}
	t2 := &TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}

	assert.Equal(t, false, isSameTree(t1, t2))
}
