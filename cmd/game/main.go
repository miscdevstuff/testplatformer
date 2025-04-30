package main

import (
	"log"
	"testplatformer/internal/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Test Platformer")
	if err := ebiten.RunGame(scenes.NewGame()); err != nil {
		log.Fatal(err)
	}
}