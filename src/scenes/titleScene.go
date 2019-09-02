package scenes

import (
	"font/common"
	"font/common/constant"
	"font/sprites"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"image/color"
)

type Title struct {
}

func (t *Title) Update(screen *ebiten.Image) {
	text.Draw(screen, common.TitleMenuString, common.FontObject.FontHuakang, constant.ScreenWidth/2-(5*12), constant.ScreenHeight-200, color.White)

	op := &ebiten.DrawImageOptions{}
	hanafudaTitleW, hanafudaTitleH := sprites.TitleImage.Size()
	op.GeoM.Translate(float64(constant.ScreenWidth/2)-float64(hanafudaTitleW/2), float64(hanafudaTitleH)/2.0)
	screen.DrawImage(sprites.TitleImage, op)

	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		constant.GameMode = constant.ModeGame
	}
}
