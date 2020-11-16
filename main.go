package main

import (
	"log"

	_ "image/png"

	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/basic"
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/game"
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/resources"
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/strategy"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func main() {
	fieldSize := 64.0
	gameSizeX := 1024
	gameSizeY := 768
	warriorImage, _, _ := ebitenutil.NewImageFromFile("data/gfx/warrior.png")

	mainGame := &game.Game{}
	creator := strategy.InitPlayerCreator()
	mainPlayer, playerID := creator.CreatePlayer("player")
	mainGame.InitGame(fieldSize)
	mainGame.AddResource(warriorImage, "warrior", resources.Graphics)
	mainGame.AddPlayer(mainPlayer)

	entity1 := strategy.InitStrategyEntity(1, 1, "warrior", playerID)
	entity2 := strategy.InitStrategyEntity(3, 2, "warrior", playerID)

	definiton := &game.BoardDefinition{NumberOfColumns: 12, NumberOfRows: 10, FieldSize: int(fieldSize)}
	board := &game.Board{}
	board.SetBoardDefinition(definiton)
	mainGame.SetBoard(board)

	mainGame.AppendEntity(entity1, basic.StrategyEntity)
	mainGame.AppendEntity(entity2, basic.StrategyEntity)

	ebiten.SetWindowSize(gameSizeX, gameSizeY)
	ebiten.SetWindowTitle("BattleMaster Tactics")
	if err := ebiten.RunGame(mainGame); err != nil {
		log.Fatal(err)
	}
}
