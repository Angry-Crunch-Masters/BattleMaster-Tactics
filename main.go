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
	fieldSize := 32.0
	gameSizeX := 640
	gameSizeY := 640
	zoom := 0.5
	warriorImage, _, _ := ebitenutil.NewImageFromFile("data/gfx/warrior.png")
	grassImage, _, _ := ebitenutil.NewImageFromFile("data/gfx/grass.png")

	mainGame := &game.Game{}
	creator := strategy.InitPlayerCreator()
	mainPlayer, playerID := creator.CreatePlayer("player")
	anotherPlayer, anotherPlayerID := creator.CreatePlayer("player")
	mainGame.InitGame(fieldSize, zoom)
	mainGame.AddResource(warriorImage, "warrior", resources.Graphics)
	mainGame.AddResource(grassImage, "grass", resources.Graphics)
	mainGame.AddPlayer(mainPlayer)
	mainGame.AddPlayer(anotherPlayer)

	entity1 := strategy.InitStrategyEntity(1, 1, "warrior", playerID)
	entity2 := strategy.InitStrategyEntity(3, 2, "warrior", playerID)
	entity3 := strategy.InitStrategyEntity(6, 4, "warrior", anotherPlayerID)

	definiton := &game.BoardDefinition{NumberOfColumns: 8, NumberOfRows: 8, FieldSize: int(fieldSize)}
	board := &game.Board{}
	board.SetBoardDefinition(definiton)
	mainGame.SetBoard(board)

	mainGame.AppendEntity(entity1, basic.StrategyEntity)
	mainGame.AppendEntity(entity2, basic.StrategyEntity)
	mainGame.AppendEntity(entity3, basic.StrategyEntity)

	ebiten.SetWindowSize(gameSizeX, gameSizeY)
	ebiten.SetWindowTitle("BattleMaster Tactics")
	if err := ebiten.RunGame(mainGame); err != nil {
		log.Fatal(err)
	}
}
