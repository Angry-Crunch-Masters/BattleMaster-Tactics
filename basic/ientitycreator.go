package basic

//IEntityCreator is interface for creating entities
type IEntityCreator interface {
	CreateEntity(inputData EntityBasicData) IEntity
}
