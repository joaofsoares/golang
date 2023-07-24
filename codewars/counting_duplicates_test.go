package codewars

import "testing"

func TestCountingDuplicates1(t *testing.T) {

	input := "abcde"

	expected := 0

	result := DuplicateCount(input)

	if expected != result {
		t.Fatalf("Expected = %d, is not equals result = %d", expected, result)
	}

}

func TestCountingDuplicates2(t *testing.T) {

	input := "aabbcde"

	expected := 2

	result := DuplicateCount(input)

	if expected != result {
		t.Fatalf("Expected = %d, is not equals result = %d", expected, result)
	}

}
