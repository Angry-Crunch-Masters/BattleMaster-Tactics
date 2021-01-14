package game

import (
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/basic"
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/graphics"
	"github.com/Angry-Crunch-Masters/BattleMaster-Tactics/resources"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

//Game is struct defining all game logic
type Game struct {
	board                        *Board
	resourcesManager             *resources.ResourceManager
	entitiesManager              *basic.EntityManager
	entities                     []basic.IEntity
	state                        InputState
	cameraXOffset, cameraYOffset float64
	canvasCameraOffset           float64
	zoom                         float64
	players                      []basic.IPlayer
	mainPlayer                   basic.IPlayer
}

//InitGame is used to init game
func (game *Game) InitGame(canvasCameraOffset float64, zoom float64) {
	game.resourcesManager = &resources.ResourceManager{}
	game.entitiesManager = basic.CreateEntityManager()
	game.players = make([]basic.IPlayer, 0)
	game.resourcesManager.InitResourceManager()
	game.cameraYOffset = 0
	game.cameraXOffset = 0
	game.canvasCameraOffset = canvasCameraOffset
	game.zoom = zoom
}

//AddPlayer is used to add player for game
func (game *Game) AddPlayer(player basic.IPlayer, isMainPlayer bool) {
	game.players = append(game.players, player)
	if isMainPlayer {
		game.mainPlayer = player
	}
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
	game.getUnitInfo()
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

func (game *Game) getUnitInfo() {
	if game.mainPlayer != nil {
		if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			game.mainPlayer.PreviousEntity()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyE) {
			game.mainPlayer.NextEntity()
		}
	}
}

//Draw is used to draw data on screen
func (game *Game) Draw(screen *ebiten.Image) {
	if game.board != nil {
		mainCanvas := &graphics.Canvas{}
		mainCanvas.InitCanvas(screen)
		mainCanvas.SetCameraOffset(0, 0)

		gameImage := ebiten.NewImage(256, 256)
		gameCanvas := &graphics.Canvas{}
		gameCanvas.InitCanvas(gameImage)
		gameCanvas.SetCameraOffset(game.cameraXOffset, game.cameraYOffset)
		game.board.DrawBoard(gameCanvas, game.resourcesManager, *game.entitiesManager.GetEntities())
		mainCanvas.DrawCanvas(gameCanvas, 32, 32)
	}
}

//AddCreators adds creators for game
func (game *Game) AddCreators(creators map[basic.EntityType]basic.IEntityCreator) {
	for typeKey, element := range creators {
		game.entitiesManager.AddEntityCreator(element, typeKey)
	}
}

//Layout is used to return layout size
func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int((float64)(outsideWidth) * game.zoom), int((float64)(outsideHeight) * game.zoom)
}

//AddResource is used to add resource to resourcesManager
func (game *Game) AddResource(item interface{}, name string, resourceType resources.ResourceType) {
	game.resourcesManager.AddResource(item, name, resourceType)
}

//AppendEntity adds entity to game pool
func (game *Game) AppendEntity(entity basic.IEntity, entityType basic.EntityType) {
	game.entitiesManager.AddExistingEntity(entity)
}

//CreateEntity creates entity for game
func (game *Game) CreateEntity(inputData basic.EntityBasicData, entityType basic.EntityType) {
	game.mainPlayer.AddEntity(game.entitiesManager.AddEntity(entityType, inputData))
}
