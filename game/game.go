package game

import (
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/graphics"
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/resources"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

//Game is struct defining all game logic
type Game struct {
	board                        *Board
	manager                      *resources.ResourceManager
	entities                     *IEntity
	state                        InputState
	cameraXOffset, cameraYOffset float64
	canvasCameraOffset           float64
}

//InitGame is used to init game
func (game *Game) InitGame(canvasCameraOffset float64) {
	game.manager = &resources.ResourceManager{}
	game.manager.InitResourceManager()
	game.cameraYOffset = 0
	game.cameraXOffset = 0
	game.canvasCameraOffset = canvasCameraOffset
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
		MouseY:           y,
	}
	game.getCameraOffsetInfo()
	return nil
}

func (game *Game) getCameraOffsetInfo() {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		game.cameraXOffset += game.canvasCameraOffset
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		game.cameraXOffset -= game.canvasCameraOffset
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		game.cameraYOffset += game.canvasCameraOffset
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		game.cameraYOffset -= game.canvasCameraOffset
	}
}

//Draw is used to draw data on screen
func (game *Game) Draw(screen *ebiten.Image) {
	if game.board != nil {
		mainCanvas := &graphics.Canvas{}
		mainCanvas.InitCanvas(screen)
		mainCanvas.SetCameraOffset(game.cameraXOffset, game.cameraYOffset)
		game.board.DrawBoard(mainCanvas, game.manager)
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
