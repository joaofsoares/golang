package codewars

import (
	"learn/helper"
	"testing"
)

func TestSmallerThanMe1(t *testing.T) {

	input := []int{5, 4, 3, 2, 1}

	expected := []int{4, 3, 2, 1, 0}

	result := Smaller(input)

	if !helper.AreEquals(expected, result) {
		t.Fatalf("expected = %v is not equals result = %v", expected, result)
	}
}

func TestSmallerThanMe2(t *testing.T) {

	input := []int{1, 2, 0}

	expected := []int{1, 1, 0}

	result := Smaller(input)

	if !helper.AreEquals(expected, result) {
		t.Fatalf("expected = %v is not equals result = %v", expected, result)
	}
}
