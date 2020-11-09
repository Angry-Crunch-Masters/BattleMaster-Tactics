package game

import (
	"errors"
	"image/color"

	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/graphics"
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/resources"
	"github.com/hajimehoshi/ebiten"
)

//Board is struct used to implement board for game, which uses BoardDefinition
type Board struct {
	definition *BoardDefinition
	entities   []IEntity
}

//DrawBoard is used to draw board
func (board *Board) DrawBoard(canvas graphics.ICanvas, manager *resources.ResourceManager) error {
	if board.definition != nil {
		err := board.drawBoard(canvas)
		if err != nil {
			return err
		}
		for _, element := range board.entities {
			graphics := manager.GetResource(element.GetResourceName(), resources.Graphics)
			if graphics != nil {
				asset, ok := graphics.Object.(*ebiten.Image)
				if ok && err == nil {
					err = board.drawElement(canvas, element, asset)
				} else {
					return err
				}
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
func (board *Board) drawBoard(canvas graphics.ICanvas) error {
	var fieldColor color.Color
	fieldColor = color.RGBA{0x47, 0x84, 0x28, 0xFF}
	//surface.Fill(fieldColor)
	if board.definition.NumberOfColumns > 0 && board.definition.NumberOfRows > 0 && board.definition.FieldSize > 0 {
		for x := 0; x < board.definition.NumberOfColumns; x++ {
			for y := 0; y < board.definition.NumberOfRows; y++ {
				if (x+y)%2 == 1 {
					fieldColor = color.RGBA{0x7F, 0x59, 0xB6, 0xFF}
				} else {
					fieldColor = color.RGBA{0x2A, 0x55, 0x6D, 0xFF}
				}
				canvas.DrawRect(float64(x*board.definition.FieldSize), float64(y*board.definition.FieldSize),
					float64(board.definition.FieldSize), float64(board.definition.FieldSize), fieldColor)
			}
		}
	} else {
		return errors.New("BoardDefinition: bad board definition")
	}
	return nil
}

//DrawElement is used to draw element on surface, using rules from definition
func (board *Board) drawElement(canvas graphics.ICanvas, entity IEntity, graphicsResource *ebiten.Image) error {
	canvas.DrawImage(graphicsResource, float64(entity.GetX()*board.definition.FieldSize), float64(entity.GetY()*board.definition.FieldSize))
	return nil
}
