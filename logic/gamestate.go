package logic

//GameState defines state of game
type GameState struct {
	turnNumber int
	turnReady  bool
}

//CreateGameState is constructor for GameState object
func CreateGameState() *GameState {
	return &GameState{turnNumber: 0, turnReady: false}
}

//GetTurnNumber gets turn number
func (state *GameState) GetTurnNumber() int {
	return state.turnNumber
}

//NextTurn invokes next turn
func (state *GameState) NextTurn() {
	state.turnReady = false
	state.turnNumber++
}

//SetTurnReady set actual turn as ready
func (state *GameState) SetTurnReady() {
	state.turnReady = true
}

//Check if turn is ready
func (state *GameState) IsTurnReady() bool {
	return state.turnReady
}
