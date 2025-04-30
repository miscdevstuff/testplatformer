package scenes

import (
	"image/color"
	"testplatformer/internal/entities"
	"testplatformer/internal/systems"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

// Game represents the main game state
type Game struct {
	player     *entities.Player
	input      *systems.InputSystem
	platforms  []*entities.Platform
	debugMode  bool
}

// NewGame creates a new game instance
func NewGame() *Game {
	// Create input system
	input := systems.NewInputSystem(ScreenWidth, ScreenHeight)
	
	// Create player entity
	player := entities.NewPlayer(100, 100)
	
	// Create platforms
	platforms := []*entities.Platform{
		entities.NewPlatform(0, ScreenHeight-50, ScreenWidth, 50, color.RGBA{0, 100, 0, 255}),
		entities.NewPlatform(100, 300, 200, 20, color.RGBA{100, 50, 0, 255}),
		entities.NewPlatform(400, 250, 150, 20, color.RGBA{100, 50, 0, 255}),
	}
	
	return &Game{
		player:    player,
		input:     input,
		platforms: platforms,
		debugMode: true,
	}
}

// Update updates the game state
func (g *Game) Update() error {
	// Update input
	g.input.Update()
	
	// Update player with input
	g.player.Update(g.input)
	
	// Check platform collisions
	g.player.IsOnGround = false
	for _, platform := range g.platforms {
		if g.player.CheckPlatformCollision(platform) {
			g.player.IsOnGround = true
		}
	}
	
	// Toggle debug mode
	if ebiten.IsKeyPressed(ebiten.KeyF1) {
		g.debugMode = !g.debugMode
	}
	
	return nil
}

// Draw draws the game
func (g *Game) Draw(screen *ebiten.Image) {
	// Fill background
	screen.Fill(color.RGBA{128, 160, 255, 255})
	
	// Draw platforms
	for _, platform := range g.platforms {
		platform.Draw(screen)
	}
	
	// Draw player
	g.player.Draw(screen)
	
	// Draw debug info
	if g.debugMode {
		ebitenutil.DebugPrint(screen, "Use arrow keys to move, space to jump")
	}
}

// Layout defines the game's logical resolution
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}