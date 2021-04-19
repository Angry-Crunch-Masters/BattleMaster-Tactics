package strategy

import "github.com/Angry-Crunch-Masters/BattleMaster-Tactics/basic"

//IStrategyEntity defines strategy entity (warrior, horse, catapult, etc.)
type IStrategyEntity interface {
	basic.IEntity
	AddPossibleAction(actionName string)
}
