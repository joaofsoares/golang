package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayGetName(t *testing.T) {

	p := &Player{Name: "Foo", Age: 42}

	assert.Equal(t, "Foo", p.Name)
	assert.Equal(t, 42, p.Age)
}

func TestGameMaster(t *testing.T) {
	var gm GameMaster

	gm.SetName("Bar")
	gm.SetGame("RPG")

	assert.Equal(t, "Bar", gm.GetName())
	assert.Equal(t, "RPG", gm.GetGame())

	newGm := &GameMaster{name: "foo", game: "WorldOfDarkness"}

	assert.Equal(t, "foo", newGm.name)
	assert.Equal(t, "foo", newGm.GetName())
	assert.Equal(t, "WorldOfDarkness", newGm.game)
	assert.Equal(t, "WorldOfDarkness", newGm.GetGame())
}
