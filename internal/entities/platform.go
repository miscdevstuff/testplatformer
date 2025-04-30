package entities

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Platform represents a solid platform in the game world
type Platform struct {
	X, Y          float64
	Width, Height float64
	Color         color.RGBA
}

// NewPlatform creates a new platform entity
func NewPlatform(x, y, width, height float64, color color.RGBA) *Platform {
	return &Platform{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Color:  color,
	}
}

// Draw draws the platform
func (p *Platform) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.X, p.Y, p.Width, p.Height, p.Color)
}