package interfaces

import (
	"fmt"
	"math/rand"
)

type Player interface {
	RollDices() int
	PassTurn()
}

type Fighter struct {
	level int
}

func (f *Fighter) getLevel() int {
	return f.level
}

func (f *Fighter) RollDices() int {
	return rand.Intn(20)
}

func (f *Fighter) PassTurn() {
	fmt.Println("No action taken")
}
