package game

//IEntity defines interface for entities
type IEntity interface {
	GetX() int
	GetY() int
	GetResourceName() string
}
