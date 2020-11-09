package strategy

//IStrategyEntity defines strategy entity (warrior, horse, catapult, etc.)
type IStrategyEntity interface {
	GetPlayer() IPlayer
}
