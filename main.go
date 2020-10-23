package main

import (
	"log"

	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/game"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	entity1 := &game.BoardEntity{X: 0, Y: 0}
	entity2 := &game.BoardEntity{X: 5, Y: 1}
	mainGame := &game.Game{}
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
