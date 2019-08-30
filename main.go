package main

import (
	"font/src"
	"github.com/hajimehoshi/ebiten"
	"log"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

func main() {
	g := src.NewGame()

	if err := ebiten.Run(g.Update, screenWidth, screenHeight, 1, "Font (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
