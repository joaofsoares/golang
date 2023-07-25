package codewars

import "testing"

func TestFindMiddleElement1(t *testing.T) {

	input := [3]int{2, 3, 1}
	expected := 0

	res := Gimme(input)

	if expected != res {
		t.Fatalf("Expected %d, is not equals to Result = %d", expected, res)
	}

}

func TestFindMiddleElement2(t *testing.T) {

	input := [3]int{5, 10, 14}
	expected := 1

	res := Gimme(input)

	if expected != res {
		t.Fatalf("Expected %d, is not equals to Result = %d", expected, res)
	}

}
