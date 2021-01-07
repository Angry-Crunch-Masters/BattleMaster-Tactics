//Package basic is used for basic structures
package basic

//BasicEntity is structure for all basic entites
type BasicEntity struct {
	id       uint
	x, y     int
	resource string
}

//GetID is used to get id of entity
func (entity *BasicEntity) GetID() uint {
	return entity.id
}

//GetX is used to get X coordinate of entity
func (entity *BasicEntity) GetX() int {
	return entity.x
}

//GetY is used to get Y coordinate of entity
func (entity *BasicEntity) GetY() int {
	return entity.y
}

//GetResourceName is used to get name of resource for entity
func (entity *BasicEntity) GetResourceName() string {
	return entity.resource
}

//InitEntity inits basic entity
func InitEntity(x int, y int, name string) *BasicEntity {
	entity := &BasicEntity{}
	entity.x = x
	entity.y = y
	entity.resource = name
	return entity
}
