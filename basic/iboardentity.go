package basic

//IEntity defines interface for entities
type IEntity interface {
	GetID() uint
	GetX() int
	GetY() int
	GetResourceName() string
}
