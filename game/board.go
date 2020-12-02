package game

import (
	"errors"

	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/basic"
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/graphics"
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/resources"
	"github.com/hajimehoshi/ebiten"
)

//Board is struct used to implement board for game, which uses BoardDefinition
type Board struct {
	definition *BoardDefinition
	entities   []basic.IEntity
}

//DrawBoard is used to draw board
func (board *Board) DrawBoard(canvas graphics.ICanvas, manager *resources.ResourceManager) error {
	if board.definition != nil {
		err := board.drawBoard(canvas, manager)
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
func (board *Board) AppendEntity(entity basic.IEntity) {
	if board.entities == nil {
		board.entities = []basic.IEntity{}
	}
	if entity != nil {
		board.entities = append(board.entities, entity)
	}
}

//DrawBoard is used to draw combat board
func (board *Board) drawBoard(canvas graphics.ICanvas, manager *resources.ResourceManager) error {
	graphics := manager.GetResource("grass", resources.Graphics)
	if graphics == nil {
		return errors.New("")
	}
	if board.definition.NumberOfColumns > 0 && board.definition.NumberOfRows > 0 && board.definition.FieldSize > 0 {
		for x := 0; x < board.definition.NumberOfColumns; x++ {
			for y := 0; y < board.definition.NumberOfRows; y++ {
				canvas.DrawImage(graphics.Object.(*ebiten.Image), float64(x*board.definition.FieldSize), float64(y*board.definition.FieldSize))
			}
		}
	} else {
		return errors.New("BoardDefinition: bad board definition")
	}
	return nil
}

//DrawElement is used to draw element on surface, using rules from definition
func (board *Board) drawElement(canvas graphics.ICanvas, entity basic.IEntity, graphicsResource *ebiten.Image) error {
	canvas.DrawImage(graphicsResource, float64(entity.GetX()*board.definition.FieldSize), float64(entity.GetY()*board.definition.FieldSize))
	return nil
}
