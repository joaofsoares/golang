package codewars

import "testing"

func TestRakeGarden(t *testing.T) {
	garden := "slug spider rock gravel gravel gravel gravel gravel gravel gravel"

	expected := "gravel gravel rock gravel gravel gravel gravel gravel gravel gravel"

	result := RakeGarden(garden)

	if expected != result {
		t.Fatalf("expected = %s, is not equals to %s", expected, result)
	}

}
