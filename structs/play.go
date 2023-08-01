package structs

type Player struct {
	Name string
	Age  int
}

type GameMaster struct {
	name string
	game string
}

func (gm *GameMaster) SetName(name string) {
	gm.name = name
}

func (gm *GameMaster) SetGame(game string) {
	gm.game = game
}

func (gm *GameMaster) GetName() string {
	return gm.name
}

func (gm *GameMaster) GetGame() string {
	return gm.game
}
