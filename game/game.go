package game

import (
	"github.com/hajimehoshi/ebiten"
)

//Game is struct defining all game logic
type Game struct {
	Definition *BoardDefinition
}

//Update is used to update logic
func (g *Game) Update() error {
	return nil
}

//Draw is used to draw data on screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.Definition.DrawBoard(screen)
}

//Layout is used to return layout size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
