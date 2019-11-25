/**
 * @Author: Tomonori
 * @Date: 2019/11/12 11:22
 * @Desc:
 */
package components

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"hanafuda/common"
	"image/color"
)

type Button struct {
	Title           string
	Size            string
	Type            string //black or white or transparent-white or transparent-black
	btnSkin         *skin
	texture         *ebiten.Image
	bgTexture       *ebiten.Image
	tx              float64
	ty              float64
	Width           int
	Height          int
	titleFont       font.Face
	titleFontHeight int
}

type skin struct {
	Color      color.Color
	Background color.Color
	TitleColor color.Color
}

func NewButton(title, size, btnType string) *Button {
	button := &Button{
		Title: title,
		Size:  size,
		Type:  btnType,
	}
	button.selectBtnSize()
	button.selectBtnType()
	return button
}

func (b *Button) Translate(tx, ty float64) {
	b.tx, b.ty = tx, ty
}

func (b *Button) Persist(screen *ebiten.Image) {
	b.update(screen)
}

func (b *Button) Click(fun func()) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		fx := float64(x)
		fy := float64(y)
		if b.tx <= fx && fx < b.tx+float64(b.Width) && b.ty <= fy && fy < b.ty+float64(b.Height) {
			fun()
		}
	}
}

func (b *Button) update(screen *ebiten.Image) {
	b.texture, _ = ebiten.NewImage(b.Width, b.Height, ebiten.FilterNearest)
	b.texture.Fill(b.btnSkin.Color)

	b.bgTexture, _ = ebiten.NewImage(b.Width-10, b.Height-10, ebiten.FilterNearest)
	btnBgW, btnBgH := b.bgTexture.Size()
	b.bgTexture.Fill(b.btnSkin.Background)
	bgOpt := &ebiten.DrawImageOptions{}
	bgOpt.GeoM.Translate(float64((b.Width-btnBgW)/2), float64((b.Height-btnBgH)/2))
	b.texture.DrawImage(b.bgTexture, bgOpt)

	textureOpt := &ebiten.DrawImageOptions{}
	textureOpt.GeoM.Translate(b.tx, b.ty)

	screen.DrawImage(b.texture, textureOpt)

	text.Draw(screen, b.Title, b.titleFont, int(b.tx)+b.Width/2/2, int(b.ty)+b.titleFontHeight+b.titleFontHeight/2, b.btnSkin.TitleColor)
}

func (b *Button) selectBtnSize() {
	switch b.Size {
	case "small":
		b.Width, b.Height = 100, 32
		b.titleFont = common.FontObject.FontFiraCodeM
		b.titleFontHeight = 12
	case "normal":
		b.Width, b.Height = 100, 32
		b.titleFont = common.FontObject.FontFiraCodeM
		b.titleFontHeight = 12
	case "large":
		b.Width, b.Height = 180, 55
		b.titleFont = common.FontObject.FontFiraCode
		b.titleFontHeight = 24
	default:
		b.Width, b.Height = 100, 32
		b.titleFont = common.FontObject.FontFiraCodeM
		b.titleFontHeight = 12
	}
}

func (b *Button) selectBtnType() {
	switch b.Type {
	case "black":
		b.btnSkin = &skin{
			Color:      color.Black,
			Background: color.NRGBA{R: 139, G: 137, B: 112, A: 0xff},
			TitleColor: color.White,
		}
	case "white":
		b.btnSkin = &skin{
			Color:      color.White,
			Background: color.NRGBA{R: 139, G: 137, B: 112, A: 0xff},
			TitleColor: color.White,
		}
	case "transparent-white":
		b.btnSkin = &skin{
			Color:      color.Transparent,
			Background: color.Transparent,
			TitleColor: color.White,
		}
	case "transparent-black":
		b.btnSkin = &skin{
			Color:      color.Transparent,
			Background: color.Transparent,
			TitleColor: color.Black,
		}
	}
}
