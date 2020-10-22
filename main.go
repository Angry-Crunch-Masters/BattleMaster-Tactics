package main

import (
	"log"

	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/game"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&game.Game{Definition: &game.BoardDefinition{NumberOfColumns: 4, NumberOfRows: 4, FieldSize: 48}}); err != nil {
		log.Fatal(err)
	}
}
