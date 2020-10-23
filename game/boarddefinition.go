package game

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//BoardDefinition is used to define board for simulation
type BoardDefinition struct {
	NumberOfRows    int
	NumberOfColumns int
	FieldSize       int
}

//DrawBoard is used to draw combat board
func (definition *BoardDefinition) DrawBoard(surface *ebiten.Image) error {
	var fieldColor color.Color
	if definition.NumberOfColumns > 0 && definition.NumberOfRows > 0 && definition.FieldSize > 0 {
		for x := 0; x < definition.NumberOfColumns; x++ {
			for y := 0; y < definition.NumberOfRows; y++ {
				if (x+y)%2 == 1 {
					fieldColor = color.RGBA{0xAA, 0x10, 0x10, 0xFF}
				} else {
					fieldColor = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
				}
				ebitenutil.DrawRect(surface, float64(x*definition.FieldSize), float64(y*definition.FieldSize),
					float64(definition.FieldSize), float64(definition.FieldSize), fieldColor)
			}
		}
	} else {
		return errors.New("BoardDefinition: bad board definition")
	}
	return nil
}

//DrawElement is used to draw element on surface, using rules from definition
func (definition *BoardDefinition) DrawElement(surface *ebiten.Image, entity *BoardEntity, resource string) error {
	fieldColor := color.RGBA{0x20, 0xF0, 0x55, 0xFF}
	ebitenutil.DrawRect(surface, float64(entity.X*definition.FieldSize), float64(entity.Y*definition.FieldSize),
		float64(definition.FieldSize), float64(definition.FieldSize), fieldColor)
	return nil
}
