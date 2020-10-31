package game

//IBoardEntity defines interface for entities
type IBoardEntity interface {
	GetX() int
	GetY() int
	GetResourceName() string
}
