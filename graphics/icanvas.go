package graphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

//ICanvas is used to implement basic canvas system used for rendering
type ICanvas interface {
	DrawRect(x, y, width, height float64, clr color.Color)
	Fill(clr color.Color)
	DrawImage(img *ebiten.Image, options *ebiten.DrawImageOptions)
}
