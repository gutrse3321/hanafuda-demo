package main

import (
	"font/common/constant"
	"font/src"
	"github.com/hajimehoshi/ebiten"
	"log"
)

func main() {
	g := src.NewGame()

	if err := ebiten.Run(g.Update, constant.ScreenWidth, constant.ScreenHeight, 1, "Hanafuda (花札 Demo)"); err != nil {
		log.Fatal(err)
	}
}
