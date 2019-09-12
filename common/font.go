package common

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"hanafuda/resources/fonts"
	"io/ioutil"
	"log"
)

var FontObject *Font

type Font struct {
	FontFiraCodeBig font.Face
	FontHuakangBig  font.Face

	FontFiraCode font.Face
	FontHuakang  font.Face

	FontHuakangM  font.Face
	FontFiraCodeM font.Face
}

var (
	TitleMenuString = "Press Enter to Start"
)

func (f *Font) Init() {
	tt, err := truetype.Parse(fonts.Font_firacode)
	if err != nil {
		log.Fatal(err)
	}

	huakangFile, _ := ioutil.ReadFile("resources/fonts/huakang.ttf")
	ttH, err := freetype.ParseFont(huakangFile)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	/**
	big size
	*/
	f.FontFiraCodeBig = truetype.NewFace(tt, &truetype.Options{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	f.FontHuakangBig = truetype.NewFace(ttH, &truetype.Options{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	/**
	normal size
	*/
	f.FontFiraCode = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	f.FontHuakang = truetype.NewFace(ttH, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	/**
	M
	small size
	*/
	f.FontFiraCodeM = truetype.NewFace(tt, &truetype.Options{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	f.FontHuakangM = truetype.NewFace(ttH, &truetype.Options{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	FontObject = f
}
