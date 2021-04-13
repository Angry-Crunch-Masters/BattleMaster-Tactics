package strategy

import "github.com/Angry-Crunch-Masters/BattleMaster-Tactics/basic"

//BasicStrategyEntity is used for basic strategy entities like wariior (etc.)
type BasicStrategyEntity struct {
	*basic.BasicEntity
	ownerID          int
	availableActions []IAction
}

//GetOwner gets owner of entity
func (entity *BasicStrategyEntity) GetOwner() int {
	return entity.ownerID
}

func (entity *BasicStrategyEntity) AddPossibleAction(action IAction) {
	entity.availableActions = append(entity.availableActions, action)
}

//InitStrategyEntity inits strategy entity
func InitStrategyEntity(x int, y int, name string, ownerID int) *BasicStrategyEntity {
	entity := &BasicStrategyEntity{}
	entity.BasicEntity = basic.InitEntity(x, y, name)
	entity.ownerID = ownerID
	entity.availableActions = make([]IAction, 0)
	return entity
}
