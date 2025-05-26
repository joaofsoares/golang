package sum_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"learn/algorithms/sum"
)

func TestSumArray(t *testing.T) {
	tests := map[int]struct {
		input    []int
		expected int
	}{
		1: {
			input:    []int{1, 3, 5, 7, 9},
			expected: 25,
		},
	}

	for _, elem := range tests {
		assert.Equal(t, sum.SumArray(&elem.input, len(elem.input)), elem.expected, "sum array should be equals expected (sum of all array elements)")
	}

}
