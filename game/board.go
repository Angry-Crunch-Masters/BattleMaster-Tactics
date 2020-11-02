package game

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//Board is struct used to implement board for game, which uses BoardDefinition
type Board struct {
	definition *BoardDefinition
	entities   []IEntity
}

//DrawBoard is used to draw board
func (board *Board) DrawBoard(surface *ebiten.Image) error {
	if board.definition != nil {
		err := board.drawBoard(surface)
		if err != nil {
			return err
		}
		for _, element := range board.entities {
			err = board.drawElement(surface, element, "warrior")
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
func (board *Board) AppendEntity(entity IEntity) {
	if board.entities == nil {
		board.entities = []IEntity{}
	}
	if entity != nil {
		board.entities = append(board.entities, entity)
	}
}

//DrawBoard is used to draw combat board
func (board *Board) drawBoard(surface *ebiten.Image) error {
	var fieldColor color.Color
	if board.definition.NumberOfColumns > 0 && board.definition.NumberOfRows > 0 && board.definition.FieldSize > 0 {
		for x := 0; x < board.definition.NumberOfColumns; x++ {
			for y := 0; y < board.definition.NumberOfRows; y++ {
				if (x+y)%2 == 1 {
					fieldColor = color.RGBA{0xAA, 0x10, 0x10, 0xFF}
				} else {
					fieldColor = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
				}
				ebitenutil.DrawRect(surface, float64(x*board.definition.FieldSize), float64(y*board.definition.FieldSize),
					float64(board.definition.FieldSize), float64(board.definition.FieldSize), fieldColor)
			}
		}
	} else {
		return errors.New("BoardDefinition: bad board definition")
	}
	return nil
}

//DrawElement is used to draw element on surface, using rules from definition
func (board *Board) drawElement(surface *ebiten.Image, entity IEntity, resource string) error {
	fieldColor := color.RGBA{0x20, 0xF0, 0x55, 0xFF}
	ebitenutil.DrawRect(surface, float64(entity.GetX()*board.definition.FieldSize), float64(entity.GetY()*board.definition.FieldSize),
		float64(board.definition.FieldSize), float64(board.definition.FieldSize), fieldColor)
	return nil
}
