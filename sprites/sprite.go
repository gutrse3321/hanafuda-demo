package sprites

import (
	"bytes"
	"font/resources/textures"
	"github.com/hajimehoshi/ebiten"
	"image"
	"log"
)

const (
	FudaW     = 70
	FudaH     = 113
	FudaSpace = 4
)

var (
	CardImage *ebiten.Image
)

type Sprite struct {
	ImageWidth  int
	ImageHeight int
	X           int
	Y           int
	Vx          int
	Vy          int
	Angle       int
	Tag         string
}

type Sprites struct {
	SpriteArr []*Sprite
	Num       int
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(textures.Texture_card))
	if err != nil {
		log.Fatal(err)
	}
	CardImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}
