package codewars

import (
	"learn/helper"
	"testing"
)

func TestCapitalize1(t *testing.T) {
	input := "abcdef"

	expected := []string{"AbCdEf", "aBcDeF"}

	res := Capitalize(input)

	if !helper.AreEquals(expected, res) {
		t.Fatalf("Expected = %v, is not equals to Result = %v", expected, res)
	}
}

func TestCapitalize2(t *testing.T) {
	input := "codewars"

	expected := []string{"CoDeWaRs", "cOdEwArS"}

	res := Capitalize(input)

	if !helper.AreEquals(expected, res) {
		t.Fatalf("Expected = %v, is not equals to Result = %v", expected, res)
	}
}
