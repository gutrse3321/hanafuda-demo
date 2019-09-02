package scenes

import (
	"font/common"
	"font/common/constant"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"image/color"
)

type Game struct {
}

func (t *Game) Update(screen *ebiten.Image) {
	text.Draw(screen, "Welcome Gamers！", common.FontObject.FontHuakang, constant.ScreenWidth/2-(5*12), constant.ScreenHeight/2-12, color.White)
}
