package main

import (
	"log"

	_ "image/png"

	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/game"
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/resources"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func main() {
	warriorImage, _, _ := ebitenutil.NewImageFromFile("data/gfx/warrior.png")

	entity1 := game.InitEntity(1, 1, "warrior")
	entity2 := game.InitEntity(3, 2, "warrior")
	mainGame := &game.Game{}
	mainGame.InitGame()
	mainGame.AddResource(warriorImage, "warrior", resources.Graphics)

	definiton := &game.BoardDefinition{NumberOfColumns: 4, NumberOfRows: 4, FieldSize: 64}
	board := &game.Board{}
	board.SetBoardDefinition(definiton)
	board.AppendEntity(entity1)
	board.AppendEntity(entity2)

	mainGame.SetBoard(board)
	ebiten.SetWindowSize(512, 512)
	ebiten.SetWindowTitle("BattleMaster Tactics")
	if err := ebiten.RunGame(mainGame); err != nil {
		log.Fatal(err)
	}
}
