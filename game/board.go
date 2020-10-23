package game

import (
	"errors"

	"github.com/hajimehoshi/ebiten"
)

//Board is struct used to implement board for game, which uses BoardDefinition
type Board struct {
	definition *BoardDefinition
	entities   []BoardEntity
}

//DrawBoard is used to draw board
func (board *Board) DrawBoard(surface *ebiten.Image) error {
	if board.definition != nil {
		err := board.definition.DrawBoard(surface)
		if err != nil {
			return err
		}
		for _, element := range board.entities {
			err = board.definition.DrawElement(surface, &element, "")
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("Board: no definition for board")
	}
	return nil
}

//SetBoardDefinition is used to set definition of board
func (board *Board) SetBoardDefinition(definition *BoardDefinition) {
	board.definition = definition
}

//AppendEntity is used to append entity to board
func (board *Board) AppendEntity(entity *BoardEntity) {
	if board.entities == nil {
		board.entities = []BoardEntity{}
	}
	if entity != nil {
		board.entities = append(board.entities, *entity)
	}
}
