package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"hanafuda/common"
	"hanafuda/common/constant"
	"hanafuda/sprites"
	"image/color"
)

type Title struct {
}

func (t *Title) Update(screen *ebiten.Image) {
	//version & other information
	text.Draw(screen, "v0.0.1", common.FontObject.FontFiraCodeM, 50, constant.ScreenHeight-50, color.NRGBA{220, 220, 220, 0xff})

	//start content
	text.Draw(screen, common.TitleMenuString, common.FontObject.FontFiraCode, constant.ScreenWidth/2-(10*12), constant.ScreenHeight-200, color.White)

	//logo
	op := &ebiten.DrawImageOptions{}
	hanafudaTitleW, hanafudaTitleH := sprites.TitleImage.Size()
	op.GeoM.Translate(float64(constant.ScreenWidth/2)-float64(hanafudaTitleW/2), float64(hanafudaTitleH)/2.0)
	screen.DrawImage(sprites.TitleImage, op)

	//input
	if ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsKeyPressed(ebiten.KeyKPEnter) {
		constant.GameMode = constant.ModeGame
	}
}
