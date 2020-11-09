package main

import (
	"log"

	_ "image/png"

	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/game"
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/resources"
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/strategy"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func main() {
	gameSizeX := 1024
	gameSizeY := 768
	warriorImage, _, _ := ebitenutil.NewImageFromFile("data/gfx/warrior.png")

	entity1 := game.InitEntity(1, 1, "warrior")
	entity2 := game.InitEntity(3, 2, "warrior")
	mainGame := &game.Game{}
	mainPlayer := &strategy.Player{}
	mainPlayer.InitPlayer()
	mainGame.InitGame()
	mainGame.AddResource(warriorImage, "warrior", resources.Graphics)

	definiton := &game.BoardDefinition{NumberOfColumns: 20, NumberOfRows: 20, FieldSize: 64}
	board := &game.Board{}
	board.SetBoardDefinition(definiton)
	board.AppendEntity(entity1)
	board.AppendEntity(entity2)

	mainGame.SetBoard(board)
	ebiten.SetWindowSize(gameSizeX, gameSizeY)
	ebiten.SetWindowTitle("BattleMaster Tactics")
	if err := ebiten.RunGame(mainGame); err != nil {
		log.Fatal(err)
	}
}
