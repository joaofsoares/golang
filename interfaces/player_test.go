package interfaces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayerRollDices(t *testing.T) {

	player := &Fighter{level: 1}

	dice := checkRange(player)

	assert.NotEqual(t, dice, -1)
}

func checkRange(p Player) int {

	res := p.RollDices()

	if res >= 0 && res <= 20 {
		return res
	}

	return -1
}
