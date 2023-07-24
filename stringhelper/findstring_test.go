package stringhelper

import "testing"

func TestFindString(t *testing.T) {
	strs := []string{"one", "two", "three", "four", "five"}
	expected := "four"
	str := FindString(expected, &strs)

	if expected != str {
		t.Fatalf("expected = %s, is not equals to %s", expected, str)
	}
}

func TestFindStringError(t *testing.T) {
	strs := []string{"one", "two", "three", "four", "five"}
	expected := "invalid"
	str := FindString(expected, &strs)

	if str != "" {
		t.Fatalf("var str should be empty")
	}
}
