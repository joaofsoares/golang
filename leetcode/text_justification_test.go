package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextJustify1(t *testing.T) {

	input := []string{"This", "is", "an", "example", "of", "text", "justification."}

	expected := []string{"This    is    an", "example  of text", "justification.  "}

	res := fullJustify(input, 16)

	assert.Equal(t, expected, res, "Expected is different than Result")
}

func TestTextJustify2(t *testing.T) {
	input := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	expected := []string{"What   must   be", "acknowledgment  ", "shall be        "}

	res := fullJustify(input, 16)

	assert.Equal(t, expected, res, "Expected is different than Result")
}

func TestTextJustify3(t *testing.T) {
	input := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	expected := []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "}

	res := fullJustify(input, 20)

	assert.Equal(t, expected, res, "Expected is different than Result")
}

func TestTextJustify4(t *testing.T) {
	input := []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}
	expected := []string{"ask   not   what", "your country can", "do  for  you ask", "what  you can do", "for your country"}

	res := fullJustify(input, 16)

	assert.Equal(t, expected, res, "Expected is different than Result")
}

func TestTextJustify5(t *testing.T) {
	input := []string{"what", "you", "can", "do"}
	expected := []string{"what you can do "}

	res := fullJustify(input, 16)

	assert.Equal(t, expected, res, "Expected is different than Result")
}
