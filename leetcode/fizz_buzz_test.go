package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz1(t *testing.T) {
	res := fizzBuzz(3)
	assert.Equal(t, []string{"1", "2", "Fizz"}, res)
}

func TestFizzBuzz2(t *testing.T) {
	res := fizzBuzz(5)
	assert.Equal(t, []string{"1", "2", "Fizz", "4", "Buzz"}, res)
}

func TestFizzBuzz3(t *testing.T) {
	res := fizzBuzz(15)
	assert.Equal(t, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}, res)
}
