package structs_out

import (
	"fmt"
	"learn/structs"
)

func UsingPlay() {

	p := &structs.Player{Name: "Foo", Age: 42}

	fmt.Printf("Player name = %s\n", p.Name)

	var gm structs.GameMaster
	gm.SetName("Bar")
	gm.SetGame("RPG")

	fmt.Printf("GameMaster name = %s\n", gm.GetName())

}
