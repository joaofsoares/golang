package sort_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"learn/algorithms/sort"
)

func TestInsertionSort(t *testing.T) {
	tests := map[int]struct {
		input    []int
		expected []int
	}{
		1: {
			input:    []int{7, 5, 3, 1},
			expected: []int{1, 3, 5, 7},
		},
		2: {
			input:    []int{5, 9, 2, 1, 7, 4},
			expected: []int{1, 2, 4, 5, 7, 9},
		},
		3: {
			input:    []int{9, 9, 8, 8, 5, 5, 3, 3, 2, 2, 1},
			expected: []int{1, 2, 2, 3, 3, 5, 5, 8, 8, 9, 9},
		},
	}

	for _, elem := range tests {
		sort.InsertionSort(&elem.input, len(elem.input))
		assert.Equal(t, elem.input, elem.expected, "slices should be the same")
	}
}
