package strategy

//Player struct is used for player informations
type Player struct {
	entities []IStrategyEntity
}

//InitPlayer inits player used in game
func (player *Player) InitPlayer() {
	player.entities = make([]IStrategyEntity, 0)
}

//AddPlayerStrategyEntity is used to add strategy entity (warrior, etc.) for player
func (player *Player) AddPlayerStrategyEntity(entity IStrategyEntity) {
	player.entities = append(player.entities, entity)
}
