package sprites

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"hanafuda/common/constant"
	"hanafuda/resources/textures"
	"image"
	_ "image/png"
	"log"
)

var (
	CardImage  *ebiten.Image
	TitleImage *ebiten.Image

	CardSprites map[string]card
)

type Sprites struct {
}

type card struct {
	Texture     *ebiten.Image
	ImageWidth  int
	ImageHeight int
	X           int
	Y           int
	Vx          int
	Vy          int
	Angle       int
	Tag         string
}

func (s *Sprites) Init() {
	img, _, err := image.Decode(bytes.NewReader(textures.Texture_card))
	if err != nil {
		log.Fatal(err)
	}
	CardImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	//map card & set
	x := 0
	y := 0
	count := 0
	cards := make(map[string]card)
	for {
		if count+1 > constant.CountFuda {
			break
		}
		if x == 9 {
			x = 0
			y++
		}

		sx := x*constant.Wfuda + constant.WFudaSpace*x
		sy := y*constant.Hfuda + constant.WFudaSpace*y
		fudaItem := CardImage.SubImage(image.Rect(sx, sy, sx+constant.Wfuda, sy+constant.Hfuda)).(*ebiten.Image)
		cards[constant.TagFuda[count]] = card{
			Texture: fudaItem,
			Tag:     constant.TagFuda[count],
		}

		x++
		count++
	}
	CardSprites = cards

	titleImg, _, err := image.Decode(bytes.NewReader(textures.Texture_title))
	if err != nil {
		log.Fatal(err)
	}
	TitleImage, _ = ebiten.NewImageFromImage(titleImg, ebiten.FilterDefault)
}
