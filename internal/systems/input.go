package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// InputSystem handles both keyboard and touch input
type InputSystem struct {
	// Movement direction
	Left      bool
	Right     bool
	Jump      bool
	JustJumped bool
	
	// Touch areas
	leftZone  struct{ x1, y1, x2, y2 int }
	rightZone struct{ x1, y1, x2, y2 int }
	jumpZone  struct{ x1, y1, x2, y2 int }
	
	// Previous state to detect "just pressed" events
	prevJump bool
}

// NewInputSystem creates a new input system
func NewInputSystem(screenWidth, screenHeight int) *InputSystem {
	input := &InputSystem{}
	
	// Set up virtual touch zones
	// Left third of screen for moving left
	input.leftZone.x1 = 0
	input.leftZone.y1 = 0
	input.leftZone.x2 = screenWidth / 3
	input.leftZone.y2 = screenHeight
	
	// Right third of screen for moving right
	input.rightZone.x1 = screenWidth * 2 / 3
	input.rightZone.y1 = 0
	input.rightZone.x2 = screenWidth
	input.rightZone.y2 = screenHeight
	
	// Middle third of screen for jump
	input.jumpZone.x1 = screenWidth / 3
	input.jumpZone.y1 = 0
	input.jumpZone.x2 = screenWidth * 2 / 3
	input.jumpZone.y2 = screenHeight
	
	return input
}

// Update updates the input state
func (i *InputSystem) Update() {
	// Save previous state
	i.prevJump = i.Jump
	
	// Reset current state
	i.Left = false
	i.Right = false
	i.Jump = false
	
	// Process keyboard input
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		i.Left = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		i.Right = true
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		i.Jump = true
	}
	
	// Process touch input
	for _, id := range ebiten.TouchIDs() {
		x, y := ebiten.TouchPosition(id)
		
		// Check which zone was touched
		if i.inZone(x, y, i.leftZone) {
			i.Left = true
		}
		if i.inZone(x, y, i.rightZone) {
			i.Right = true
		}
		if i.inZone(x, y, i.jumpZone) {
			i.Jump = true
		}
	}
	
	// Check for "just jumped" event
	i.JustJumped = i.Jump && !i.prevJump
}

// inZone checks if a point is within a defined zone
func (i *InputSystem) inZone(x, y int, zone struct{ x1, y1, x2, y2 int }) bool {
	return x >= zone.x1 && x <= zone.x2 && y >= zone.y1 && y <= zone.y2
}