package codewars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsbnValidation1(t *testing.T) {

	input := "1112223339"

	expected := true

	res := ValidISBN10(input)

	assert.Equal(t, expected, res)
}

func TestIsbnValidation2(t *testing.T) {

	input := "111222333"

	expected := false

	res := ValidISBN10(input)

	assert.Equal(t, expected, res)
}

func TestIsbnValidation3(t *testing.T) {

	input := "1112223339X"

	expected := false

	res := ValidISBN10(input)

	assert.Equal(t, expected, res)
}

func TestIsbnValidation4(t *testing.T) {

	input := "1234554321"

	expected := true

	res := ValidISBN10(input)

	assert.Equal(t, expected, res)
}

func TestIsbnValidation5(t *testing.T) {

	input := "1234512345"

	expected := false

	res := ValidISBN10(input)

	assert.Equal(t, expected, res)
}

func TestIsbnValidation6(t *testing.T) {

	input := "048665088X"

	expected := true

	res := ValidISBN10(input)

	assert.Equal(t, expected, res)
}

func TestIsbnValidation7(t *testing.T) {

	input := "X123456788"

	expected := false

	res := ValidISBN10(input)

	assert.Equal(t, expected, res)
}

func TestIsbnValidation8(t *testing.T) {

	input := "ABCDEFGHIJ"

	expected := false

	res := ValidISBN10(input)

	assert.Equal(t, expected, res)
}

func TestIsbnValidation9(t *testing.T) {

	input := "048665088x"

	expected := true

	res := ValidISBN10(input)

	assert.Equal(t, expected, res)
}
