package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"hanafuda/common"
	"hanafuda/common/constant"
	"image/color"
)

type Over struct {
}

func (o *Over) Update(screen *ebiten.Image) {
	text.Draw(screen, "Game Over", common.FontObject.FontFiraCode, constant.ScreenWidth/2-(9*6), constant.ScreenHeight/2-12, color.White)
}
