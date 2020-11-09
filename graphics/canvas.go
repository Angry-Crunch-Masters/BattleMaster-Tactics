package graphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//Canvas is structure used for drawing on canvas/surface
type Canvas struct {
	surface *ebiten.Image
}

//InitCanvas is used to init canvas
func (canvas *Canvas) InitCanvas(surface *ebiten.Image) {
	canvas.surface = surface
}

//DrawRect is used to draw rect on canvas
func (canvas *Canvas) DrawRect(x, y, width, height float64, clr color.Color) {
	if canvas.isInBoundary(x, y, width, height) {
		ebitenutil.DrawRect(canvas.surface, x, y, width, height, clr)
	}
}

//Fill fills canvas
func (canvas *Canvas) Fill(clr color.Color) {
	canvas.surface.Fill(clr)
}

//DrawImage draws image on canvas
func (canvas *Canvas) DrawImage(img *ebiten.Image, x, y float64) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(x, y)
	canvas.surface.DrawImage(img, options)
}

func (canvas *Canvas) isInBoundary(x, y, width, height float64) bool {
	return true
}
