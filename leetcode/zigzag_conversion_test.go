package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZigZagConversion1(t *testing.T) {
	s := "PAYPALISHIRING"
	expected := "PAHNAPLSIIGYIR"
	res := ZigZagConversion(s, 3)

	assert.Equal(t, expected, res, "Expected is different than Result")

}

func TestZigZagConversion2(t *testing.T) {
	s := "PAYPALISHIRING"
	expected := "PINALSIGYAHRPI"
	res := ZigZagConversion(s, 4)

	assert.Equal(t, expected, res, "Expected is different than Result")

}

func TestZigZagConversion3(t *testing.T) {
	s := "AB"
	expected := "AB"
	res := ZigZagConversion(s, 1)

	assert.Equal(t, expected, res, "Expected is different than Result")

}

func TestZigZagConversion4(t *testing.T) {
	s := "ABC"
	expected := "ACB"
	res := ZigZagConversion(s, 2)

	assert.Equal(t, expected, res, "Expected is different than Result")

}

func TestZigZagConversion5(t *testing.T) {
	s := "AB"
	expected := "AB"
	res := ZigZagConversion(s, 2)

	assert.Equal(t, expected, res, "Expected is different than Result")

}
