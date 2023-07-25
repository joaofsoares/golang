package codewars

import "testing"

func TestWeirdCase1(t *testing.T) {
	input := "String"
	expected := "StRiNg"

	result := ToWeirdCase(input)

	if expected != result {
		t.Fatalf("Expected = %s, is not equals Result = %s", expected, result)
	}
}

func TestWeirdCase2(t *testing.T) {
	input := "Weird string case"
	expected := "WeIrD StRiNg CaSe"

	result := ToWeirdCase(input)

	if expected != result {
		t.Fatalf("Expected = %s, is not equals Result = %s", expected, result)
	}
}
