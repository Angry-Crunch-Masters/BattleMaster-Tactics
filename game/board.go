package game

import (
	"errors"

	"github.com/hajimehoshi/ebiten"
)

//Board is struct used to implement board for game, which uses BoardDefinition
type Board struct {
	definition *BoardDefinition
}

//DrawBoard is used to draw board
func (board *Board) DrawBoard(surface *ebiten.Image) error {
	if board.definition != nil {
		return board.definition.DrawBoard(surface)
	}
	return errors.New("Board: no definition for board")
}

//SetBoardDefinition is used to set definition of board
func (board *Board) SetBoardDefinition(definition *BoardDefinition) {
	board.definition = definition
}
