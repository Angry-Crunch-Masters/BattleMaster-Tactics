package game

import (
	"github.com/hajimehoshi/ebiten"
)

//Game is struct defining all game logic
type Game struct {
	board *Board
}

//SetBoard is used to set board for game
func (game *Game) SetBoard(board *Board) {
	game.board = board
}

//Update is used to update logic
func (game *Game) Update() error {
	return nil
}

//Draw is used to draw data on screen
func (game *Game) Draw(screen *ebiten.Image) {
	if game.board != nil {
		game.board.DrawBoard(screen)
	}
}

//Layout is used to return layout size
func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
