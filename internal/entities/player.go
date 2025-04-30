package entities

import (
	"image/color"
	"testplatformer/internal/systems"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	// Player physics constants
	Gravity      = 0.5
	JumpVelocity = -12.0
	MoveSpeed    = 4.0
	MaxFallSpeed = 10.0
)

// Player represents the player character
type Player struct {
	X, Y          float64
	Width, Height float64
	VelocityX     float64
	VelocityY     float64
	IsOnGround    bool
	Color         color.RGBA
}

// NewPlayer creates a new player entity
func NewPlayer(x, y float64) *Player {
	return &Player{
		X:         x,
		Y:         y,
		Width:     32,
		Height:    48,
		VelocityX: 0,
		VelocityY: 0,
		IsOnGround: false,
		Color:     color.RGBA{255, 0, 0, 255},
	}
}

// Update updates the player state
func (p *Player) Update(input *systems.InputSystem) {
	// Handle horizontal movement
	p.VelocityX = 0
	if input.Left {
		p.VelocityX = -MoveSpeed
	}
	if input.Right {
		p.VelocityX = MoveSpeed
	}
	
	// Handle jumping
	if input.JustJumped && p.IsOnGround {
		p.VelocityY = JumpVelocity
		p.IsOnGround = false
	}
	
	// Apply gravity if not on ground
	if !p.IsOnGround {
		p.VelocityY += Gravity
		
		// Limit fall speed
		if p.VelocityY > MaxFallSpeed {
			p.VelocityY = MaxFallSpeed
		}
	} else {
		// Reset vertical velocity when on ground
		p.VelocityY = 0
	}
	
	// Update position
	p.X += p.VelocityX
	p.Y += p.VelocityY
}

// Draw draws the player
func (p *Player) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.X, p.Y, p.Width, p.Height, p.Color)
}

// CheckPlatformCollision checks if the player is colliding with a platform
func (p *Player) CheckPlatformCollision(platform *Platform) bool {
	// Get player bounds
	playerLeft := p.X
	playerRight := p.X + p.Width
	playerTop := p.Y
	playerBottom := p.Y + p.Height
	
	// Get platform bounds
	platformLeft := platform.X
	platformRight := platform.X + platform.Width
	platformTop := platform.Y
	platformBottom := platform.Y + platform.Height
	
	// Check for collision
	if playerRight > platformLeft && 
	   playerLeft < platformRight &&
	   playerBottom > platformTop &&
	   playerTop < platformBottom {
		
		// Handle collision response
		// If falling onto platform from above
		if p.VelocityY > 0 && 
		   playerBottom-p.VelocityY <= platformTop+5 {
			
			// Position player on top of platform
			p.Y = platformTop - p.Height
			return true
		}
	}
	
	return false
}