package strategy

import "github.com/Angry-Crunch-Masters/BattleMaster-Tactics/basic"

//BasicStrategyEntityCreator is creator for strategic entites
type BasicStrategyEntityCreator struct {
}

//CreateEntity creates entity using BasicStrategyEntityCreator
func (creator *BasicStrategyEntityCreator) CreateEntity(inputData basic.EntityBasicData) basic.IEntity {
	strategyAsset, ok := inputData.Data.(*AdditionalStrategyEntityData)
	if ok {
		return InitStrategyEntity(inputData.X, inputData.Y, inputData.Resource, strategyAsset.OwnerID, strategyAsset.ActionsNames)
	}
	return nil
}
