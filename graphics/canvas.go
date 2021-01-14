package graphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//Canvas is structure used for drawing on canvas/surface
type Canvas struct {
	surface          *ebiten.Image
	xOffset, yOffset float64
	x, y             float64
}

//InitCanvas is used to init canvas
func (canvas *Canvas) InitCanvas(surface *ebiten.Image) {
	canvas.surface = surface
	canvas.xOffset = 0
	canvas.yOffset = 0
	canvas.x = 0
	canvas.y = 0
}

//DrawRect is used to draw rect on canvas
func (canvas *Canvas) DrawRect(x, y, width, height float64, clr color.Color) {
	if canvas.isInBoundary(x, y, width, height) {
		ebitenutil.DrawRect(canvas.surface, canvas.xOffset+x, canvas.yOffset+y, width, height, clr)
	}
}

//Fill fills canvas
func (canvas *Canvas) Fill(clr color.Color) {
	canvas.surface.Fill(clr)
}

//DrawImage draws image on canvas
func (canvas *Canvas) DrawImage(img *ebiten.Image, x, y float64) {
	width, height := img.Size()
	if canvas.isInBoundary(x, y, float64(width), float64(height)) {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(x+canvas.xOffset, y+canvas.yOffset)
		canvas.surface.DrawImage(img, options)
	}
}

//SetCameraOffset sets camera offset for canvas
func (canvas *Canvas) SetCameraOffset(xOffset, yOffset float64) {
	canvas.xOffset = xOffset
	canvas.yOffset = yOffset
}

func (canvas *Canvas) isInBoundary(x, y, width, height float64) bool {
	canvasWidth, canvasHeight := canvas.surface.Size()
	if canvas.x < x+width+canvas.xOffset && canvas.x+float64(canvasWidth) > x+canvas.xOffset &&
		canvas.y < y+height+canvas.yOffset && canvas.y+float64(canvasHeight) > y+canvas.yOffset {
		return true
	}
	return false
}

func (canvas *Canvas) getCanvasSize() (width int, height int) {
	return canvas.surface.Size()
}

//DrawCanvas draws another canvas on top of canvas
func (canvas *Canvas) DrawCanvas(drawnCanvas *Canvas, x, y float64) {
	width, height := canvas.getCanvasSize()
	if canvas.isInBoundary(drawnCanvas.x, drawnCanvas.y, float64(width), float64(height)) {
		canvas.DrawImage(drawnCanvas.surface, x, y)
	}
}
