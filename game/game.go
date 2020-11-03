package game

import (
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/resources"
	"github.com/hajimehoshi/ebiten"
)

//Game is struct defining all game logic
type Game struct {
	board    *Board
	manager  *resources.ResourceManager
	entities *IEntity
	state    InputState
}

//InitGame is used to init game
func (game *Game) InitGame() {
	game.manager = &resources.ResourceManager{}
	game.manager.InitResourceManager()
}

//SetBoard is used to set board for game
func (game *Game) SetBoard(board *Board) {
	game.board = board
}

//Update is used to update logic
func (game *Game) Update() error {
	click := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	x, y := ebiten.CursorPosition()
	game.state = InputState{
		MouseLeftClicked: click,
		MouseX:           x,
		MouseY:           y}
	return nil
}

//Draw is used to draw data on screen
func (game *Game) Draw(screen *ebiten.Image) {
	if game.board != nil {
		game.board.DrawBoard(screen, game.manager)
	}
}

//Layout is used to return layout size
func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

//AddResource is used to add resource to manager
func (game *Game) AddResource(item interface{}, name string, resourceType resources.ResourceType) {
	game.manager.AddResource(item, name, resourceType)
}
