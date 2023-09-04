package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLicenseKeyFormatting1(t *testing.T) {
	input, k := "5F3Z-2e-9-w", 4

	expected := "5F3Z-2E9W"

	res := licenseKeyFormatting(input, k)

	assert.Equal(t, expected, res)
}

func TestLicenseKeyFormatting2(t *testing.T) {
	input, k := "2-5g-3-J", 2

	expected := "2-5G-3J"

	res := licenseKeyFormatting(input, k)

	assert.Equal(t, expected, res)
}

func TestLicenseKeyFormatting3(t *testing.T) {
	input, k := "2-4A0r7-4k", 4

	expected := "24A0-R74K"

	res := licenseKeyFormatting(input, k)

	assert.Equal(t, expected, res)
}

func TestLicenseKeyFormatting4(t *testing.T) {
	input, k := "2", 2

	expected := "2"

	res := licenseKeyFormatting(input, k)

	assert.Equal(t, expected, res)
}
