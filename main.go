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

	checkedImage, _, _ := ebitenutil.NewImageFromFile("data/gfx/checked.png")
	frameImage, _, _ := ebitenutil.NewImageFromFile("data/gfx/frame.png")

	mainGame := &game.Game{}
	creator := strategy.InitPlayerCreator()
	mainPlayer, _ := creator.CreatePlayer("player")
	mainGame.InitGame(fieldSize, zoom)
	mainGame.AddResource(warriorImage, "warrior", resources.Graphics)
	mainGame.AddResource(grassImage, "grass", resources.Graphics)
	mainGame.AddResource(checkedImage, "checked", resources.Graphics)
	mainGame.AddResource(frameImage, "frame", resources.Graphics)
	mainGame.AddPlayer(mainPlayer, true)
	mainGame.AddCreators(map[basic.EntityType]basic.IEntityCreator{basic.StrategyEntity: &strategy.BasicStrategyEntityCreator{}})

	definiton := &game.BoardDefinition{NumberOfColumns: 8, NumberOfRows: 8, FieldSize: int(fieldSize)}
	board := &game.Board{}
	board.SetBoardDefinition(definiton)
	mainGame.SetBoard(board)

	mainGame.CreateEntity(basic.EntityBasicData{X: 6, Y: 4, Resource: "warrior",
		Data: &strategy.AdditionalEntityData{OwnerID: mainPlayer.GetPlayerID()}}, basic.StrategyEntity)
	mainGame.CreateEntity(basic.EntityBasicData{X: 4, Y: 6, Resource: "warrior",
		Data: &strategy.AdditionalEntityData{OwnerID: mainPlayer.GetPlayerID()}}, basic.StrategyEntity)
	mainGame.CreateEntity(basic.EntityBasicData{X: 3, Y: 3, Resource: "warrior",
		Data: &strategy.AdditionalEntityData{OwnerID: mainPlayer.GetPlayerID()}}, basic.StrategyEntity)

	ebiten.SetWindowSize(gameSizeX, gameSizeY)
	ebiten.SetWindowTitle("BattleMaster Tactics")
	if err := ebiten.RunGame(mainGame); err != nil {
		log.Fatal(err)
	}
}
