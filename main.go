package main

import (
	"github.com/hajimehoshi/ebiten"
	"hanafuda/common/constant"
	"hanafuda/src"
	"log"
)

func main() {
	g := src.NewGame()

	if err := ebiten.Run(g.Update, constant.ScreenWidth, constant.ScreenHeight, 1, "Hanafuda (花札 Demo)"); err != nil {
		log.Fatal(err)
	}
}
