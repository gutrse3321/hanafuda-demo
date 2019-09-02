package sprites

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"hanafuda/resources/textures"
	"image"
	_ "image/png"
	"log"
)

var (
	CardImage  *ebiten.Image
	TitleImage *ebiten.Image
)

type Sprites struct {
}

type Card struct {
	ImageWidth  int
	ImageHeight int
	X           int
	Y           int
	Vx          int
	Vy          int
	Angle       int
	Tag         string
}

type Cards struct {
	SpriteArr []*Card
	Num       int
}

type Title struct {
	ImageWidth  int
	ImageHeight int
	X           int
	Y           int
	Vx          int
	Vy          int
}

func (s *Sprites) Init() {
	img, _, err := image.Decode(bytes.NewReader(textures.Texture_card))
	if err != nil {
		log.Fatal(err)
	}
	CardImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	titleImg, _, err := image.Decode(bytes.NewReader(textures.Texture_title))
	if err != nil {
		log.Fatal(err)
	}
	TitleImage, _ = ebiten.NewImageFromImage(titleImg, ebiten.FilterDefault)
}
