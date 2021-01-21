package logic

type IGameState interface {
	GetTurnNumber() int
	NextTurn()
	SetTurnReady()
	IsTurnReady() bool
}
