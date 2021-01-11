package strategy

import "github.com/Angry-Crunch-Masters/BattleMaster-Tactics/basic"

//BasicStrategyEntityCreator is creator for strategic entites
type BasicStrategyEntityCreator struct {
}

//CreateEntity creates entity using BasicStrategyEntityCreator
func (creator *BasicStrategyEntityCreator) CreateEntity(x, y int, resource string, entityData interface{}) basic.IEntity {
	strategyAsset, ok := entityData.(*AdditionalEntityData)
	if ok {
		return InitStrategyEntity(x, y, resource, strategyAsset.ownerID)
	}
	return nil
}
