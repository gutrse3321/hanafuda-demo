/**
 * @Author: Tomonori
 * @Date: 2019/11/12 11:22
 * @Desc:
 */
package components

import "github.com/hajimehoshi/ebiten"

type ButtonBase struct {
	Texture *ebiten.Image
	tx      float64
	ty      float64
}

func (b *ButtonBase) Translate(tx, ty float64) {
	b.tx, b.ty = tx, ty
}

func (b *ButtonBase) New(screen *ebiten.Image, fun ...func()) {
	b.update(screen)
}

func (b *ButtonBase) update(screen *ebiten.Image) {
}
