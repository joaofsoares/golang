package codewars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIpsBetween1(t *testing.T) {
	start, end := "10.0.0.0", "10.0.0.50"

	expected := 50

	res := IpsBetween(start, end)

	assert.Equal(t, expected, res)
}

func TestIpsBetween2(t *testing.T) {
	start, end := "10.0.0.0", "10.0.1.0"

	expected := 256

	res := IpsBetween(start, end)

	assert.Equal(t, expected, res)
}

func TestIpsBetween3(t *testing.T) {
	start, end := "20.0.0.10", "20.0.1.0"

	expected := 246

	res := IpsBetween(start, end)

	assert.Equal(t, expected, res)
}

func TestIpsBetween4(t *testing.T) {
	start, end := "170.0.0.10", "170.1.0.0"

	expected := 65526

	res := IpsBetween(start, end)

	assert.Equal(t, expected, res)
}

func TestIpsBetween5(t *testing.T) {
	start, end := "180.0.0.0", "181.0.0.0"

	expected := 16777216

	res := IpsBetween(start, end)

	assert.Equal(t, expected, res)
}
