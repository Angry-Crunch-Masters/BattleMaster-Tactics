package strategy

import "github.com/Angry-Crunch-Masters/BattleMaster-Tactics/basic"

//PlayerCreator is used to create new players
type PlayerCreator struct {
	playerCounter int
}

//InitPlayerCreator is used to init player creator
func InitPlayerCreator() *PlayerCreator {
	return &PlayerCreator{playerCounter: 0}
}

//CreatePlayer is used to create player with unique id
func (creator *PlayerCreator) CreatePlayer(name string) (basic.IPlayer, int) {
	player := &Player{}
	playerID := creator.playerCounter
	player.InitPlayer(name, playerID)
	creator.playerCounter++
	return player, playerID
}
