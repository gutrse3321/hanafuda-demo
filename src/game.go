package src

import (
	"github.com/hajimehoshi/ebiten"
	"hanafuda/common"
	"hanafuda/common/constant"
	"hanafuda/sprites"
	"hanafuda/src/scenes"
	"image/color"
)

type Game struct {
}

var (
	titleScene scenes.Title
	gameScene  scenes.Game
	overScene  scenes.Over
)

func (g *Game) init() {
	fontModel := &common.Font{}
	spritesModel := &sprites.Sprites{}
	fontModel.Init()
	spritesModel.Init()
}

func (g *Game) Update(screen *ebiten.Image) error {
	screen.Fill(color.NRGBA{R: 139, G: 137, B: 112, A: 0xff})

	switch constant.GameMode {
	case constant.ModeTitle:
		titleScene.Update(screen)
	case constant.ModeGame:
		gameScene.Update(screen)
	case constant.ModeGameOver:
		overScene.Update(screen)
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return nil
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}
