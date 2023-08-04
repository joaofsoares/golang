package codewars

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRot13One(t *testing.T) {

	input := "EBG13 rknzcyr."
	expected := "ROT13 example."

	res := Rot13(input)

	assert.Equal(t, expected, res)
}

func TestRot13Two(t *testing.T) {

	input := "How can you tell an extrovert from an\nintrovert at NSA? Va gur ryringbef,\ngur rkgebireg ybbxf ng gur BGURE thl'f fubrf."
	expected := "Ubj pna lbh gryy na rkgebireg sebz na\nvagebireg ng AFN? In the elevators,\nthe extrovert looks at the OTHER guy's shoes."

	res := Rot13(input)

	assert.Equal(t, expected, res)
}

func TestRot13Three(t *testing.T) {

	input := "123"
	expected := "123"

	res := Rot13(input)

	assert.Equal(t, expected, res)
}

func TestRot13Four(t *testing.T) {

	input := "Guvf vf npghnyyl gur svefg xngn V rire znqr. Gunaxf sbe svavfuvat vg! :)"
	expected := "This is actually the first kata I ever made. Thanks for finishing it! :)"

	res := Rot13(input)

	assert.Equal(t, expected, res)
}

func TestRot13Five(t *testing.T) {

	input := "@[`{"
	expected := "@[`{"

	res := Rot13(input)

	assert.Equal(t, expected, res)
}

func TestRot13Six(t *testing.T) {

	input := 'a'    // 97
	expected := 'n' // 110

	res := input + 13

	assert.Equal(t, expected, res)
}

func TestRot13Seven(t *testing.T) {

	input := 'n'    // 110
	expected := 'z' // 122

	res := input + 12

	assert.Equal(t, expected, res)
}

func TestRot13Eight(t *testing.T) {

	input := 'A'    // 65
	expected := 'N' // 78

	res := input + 13

	assert.Equal(t, expected, res)
}

func TestRot13Nine(t *testing.T) {

	input := 'N'    // 78
	expected := 'Z' // 90

	res := input + 12

	assert.Equal(t, expected, res)
}
