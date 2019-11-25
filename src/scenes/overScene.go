package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"hanafuda/common"
	"hanafuda/common/constant"
	"hanafuda/src/components"
	"image/color"
)

type Over struct {
	btnBox *ebiten.Image
	btnBg  *ebiten.Image
}

func (o *Over) Update(screen *ebiten.Image) {
	text.Draw(screen, "Game Over", common.FontObject.FontFiraCodeBig, constant.ScreenWidth/2-(11*12), constant.ScreenHeight/2-24, color.White)

	//restart button
	restartButton := components.NewButton(
		"Restart",
		"large",
		"black",
	)
	restartButton.Translate(float64(constant.ScreenWidth/2-restartButton.Width/2), float64(constant.ScreenHeight-200))
	restartButton.Persist(screen)

	//event
	restartButton.Click(func() {
		constant.GameMode = constant.ModeGame
	})

	if ebiten.IsDrawingSkipped() {
		return
	}
}
