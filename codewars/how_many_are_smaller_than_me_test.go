package codewars

import (
	"testing"
)

func areEquals(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}

	return true
}

func TestSmallerThanMe1(t *testing.T) {

	input := []int{5, 4, 3, 2, 1}

	expected := []int{4, 3, 2, 1, 0}

	result := Smaller(input)

	if !areEquals(expected, result) {
		t.Fatalf("expected = %v is not equals result = %v", expected, result)
	}
}

func TestSmallerThanMe2(t *testing.T) {

	input := []int{1, 2, 0}

	expected := []int{1, 1, 0}

	result := Smaller(input)

	if !areEquals(expected, result) {
		t.Fatalf("expected = %v is not equals result = %v", expected, result)
	}
}
