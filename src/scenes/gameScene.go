package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"hanafuda/common/constant"
	"hanafuda/sprites"
)

type Game struct {
}

func (t *Game) Update(screen *ebiten.Image) {
	//text.Draw(screen, "Welcome Gamers！", common.FontObject.FontHuakang, constant.ScreenWidth/2-(5*12), constant.ScreenHeight/2-12, color.White)

	//TODO test sprite
	op := &ebiten.DrawImageOptions{}
	jan0W, jan0H := sprites.CardSprites.SpriteArr[0].Texture.Size()
	op.GeoM.Translate(float64(constant.ScreenWidth/2)-float64(jan0W/2), float64(jan0H)/2.0)
	screen.DrawImage(sprites.CardSprites.SpriteArr[0].Texture, op)
}
