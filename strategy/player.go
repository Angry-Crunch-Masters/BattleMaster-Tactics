package strategy

//Player struct is used for player informations
type Player struct {
	entities []IStrategyEntity
	name     string
	id       int
}

//InitPlayer inits player used in game
func (player *Player) InitPlayer(name string, id int) {
	player.entities = make([]IStrategyEntity, 0)
	player.name = name
	player.id = id
}

//AddPlayerStrategyEntity is used to add strategy entity (warrior, etc.) for player
func (player *Player) AddPlayerStrategyEntity(entity IStrategyEntity) {
	player.entities = append(player.entities, entity)
}

//GetPlayerID is used to get id of player
func (player *Player) GetPlayerID() int {
	return player.id
}
