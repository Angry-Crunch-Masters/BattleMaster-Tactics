package basic

//IEntityCreator is interface for creating entities
type IEntityCreator interface {
	CreateEntity(x, y int, resource string, entityData interface{}) IEntity
}
