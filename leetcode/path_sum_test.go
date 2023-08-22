package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSum1(t *testing.T) {
	input, target := &TreeNode{5, &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}, nil}, &TreeNode{8, &TreeNode{13, nil, nil}, &TreeNode{4, nil, &TreeNode{1, nil, nil}}}}, 22

	res := hasPathSum(input, target)

	assert.Equal(t, true, res)
}

func TestPathSum2(t *testing.T) {
	input, target := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 5

	res := hasPathSum(input, target)

	assert.Equal(t, false, res)
}
