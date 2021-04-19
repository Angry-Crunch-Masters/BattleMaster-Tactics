package strategy

import "github.com/Angry-Crunch-Masters/BattleMaster-Tactics/basic"

//BasicStrategyEntity is used for basic strategy entities like wariior (etc.)
type BasicStrategyEntity struct {
	*basic.BasicEntity
	ownerID          int
	availableActions []string
}

func (entity *BasicStrategyEntity) AddPossibleAction(actionName string) {
	entity.availableActions = append(entity.availableActions, actionName)
}

//InitStrategyEntity inits strategy entity
func InitStrategyEntity(x int, y int, name string, ownerID int, actionsNames []string) *BasicStrategyEntity {
	entity := &BasicStrategyEntity{}
	entity.BasicEntity = basic.InitEntity(x, y, name)
	entity.ownerID = ownerID
	entity.availableActions = actionsNames
	return entity
}
