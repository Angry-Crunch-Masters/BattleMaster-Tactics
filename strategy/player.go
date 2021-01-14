package strategy

import "github.com/Angry-Crunch-Masters/BattleMaster-Tactics/basic"

//Player struct is used for player informations
type Player struct {
	entities            []IStrategyEntity
	actualEntityCounter int
	name                string
	id                  int
}

//InitPlayer inits player used in game
func (player *Player) InitPlayer(name string, id int) {
	player.entities = make([]IStrategyEntity, 0)
	player.name = name
	player.id = id
	player.actualEntityCounter = 0
}

//AddEntity is used to add strategy entity (warrior, etc.) for player
func (player *Player) AddEntity(entity basic.IEntity) {
	strategicEntity, ok := entity.(IStrategyEntity)
	if ok {
		player.entities = append(player.entities, strategicEntity)
		player.actualEntityCounter = len(player.entities) - 1
	}
}

//NextEntity sets next strategic entity
func (player *Player) NextEntity() {
	entitesNumber := len(player.entities)
	if entitesNumber != 0 {
		if entitesNumber > player.actualEntityCounter+1 {
			player.actualEntityCounter++
		} else {
			player.actualEntityCounter = 0
		}
	}
}

//PreviousEntity changes strategy entity to previous
func (player *Player) PreviousEntity() {
	entitesNumber := len(player.entities)
	if entitesNumber != 0 {
		if player.actualEntityCounter-1 < 0 {
			player.actualEntityCounter = entitesNumber - 1
		} else {
			player.actualEntityCounter--
		}
	}
}

//GetPlayerID is used to get id of player
func (player *Player) GetPlayerID() int {
	return player.id
}
