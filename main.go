package main

import (
	"log"

	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/game"
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/resources"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	manager := &resources.ResourceManager{}
	manager.InitResourceManager()
	//warriorImage, _, err := ebitenutil.NewImageFromFile("data/gfx/warrior.png")

	entity1 := game.InitEntity(1, 2, "")
	entity2 := game.InitEntity(5, 7, "")
	mainGame := &game.Game{}
	mainGame.InitGame()

	definiton := &game.BoardDefinition{NumberOfColumns: 8, NumberOfRows: 8, FieldSize: 48}
	board := &game.Board{}
	board.SetBoardDefinition(definiton)
	board.AppendEntity(entity1)
	board.AppendEntity(entity2)

	mainGame.SetBoard(board)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("BattleMaster Tactics")
	if err := ebiten.RunGame(mainGame); err != nil {
		log.Fatal(err)
	}
}
