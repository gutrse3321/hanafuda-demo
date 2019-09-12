package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"hanafuda/common"
	"hanafuda/common/constant"
	"image/color"
)

type Over struct {
	btnBox *ebiten.Image
	btnBg  *ebiten.Image
}

func (o *Over) Update(screen *ebiten.Image) {
	text.Draw(screen, "Game Over", common.FontObject.FontFiraCodeBig, constant.ScreenWidth/2-(11*12), constant.ScreenHeight/2-24, color.White)

	//restart button
	if o.btnBox == nil {
		o.btnBox, _ = ebiten.NewImage(180, 55, ebiten.FilterNearest)
	}
	o.btnBox.Fill(color.Black)
	btnBoxW, _ := o.btnBox.Size()
	if o.btnBg == nil {
		o.btnBg, _ = ebiten.NewImage(170, 45, ebiten.FilterNearest)
	}
	o.btnBg.Fill(color.NRGBA{R: 139, G: 137, B: 112, A: 0xff})
	btnBgW, _ := o.btnBg.Size()

	boxOpt := &ebiten.DrawImageOptions{}
	bgOpt := &ebiten.DrawImageOptions{}

	boxOpt.GeoM.Translate(float64(constant.ScreenWidth/2-btnBoxW/2), float64(constant.ScreenHeight-200))
	bgOpt.GeoM.Translate(float64(constant.ScreenWidth/2-btnBgW/2), float64(constant.ScreenHeight-200+5))

	screen.DrawImage(o.btnBox, boxOpt)
	screen.DrawImage(o.btnBg, bgOpt)

	text.Draw(screen, "Restart", common.FontObject.FontFiraCode, constant.ScreenWidth/2-(8*6), constant.ScreenHeight-200+35, color.White)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		constant.GameMode = constant.ModeGame
	}
	//tools.New().Ticker(5, func() {
	//	constant.GameMode = constant.ModeTitle
	//})
}
