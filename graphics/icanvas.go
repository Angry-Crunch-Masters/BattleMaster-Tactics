package graphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

//ICanvas is used to implement basic canvas system used for rendering
type ICanvas interface {
	DrawRect(x, y, width, height float64, clr color.Color)
	Fill(clr color.Color)
	DrawImage(img *ebiten.Image, x, y float64)
	DrawCanvas(drawnCanvas ICanvas, x, y float64)
	GetPosition() (float64, float64)
	getSurface() *ebiten.Image
}
