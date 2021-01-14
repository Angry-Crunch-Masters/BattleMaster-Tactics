package basic

//IPlayer defines player interface
type IPlayer interface {
	GetPlayerID() int
	PreviousEntity()
	NextEntity()
	AddEntity(entity IEntity)
}
