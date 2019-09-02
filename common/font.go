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
	FontFiraCode font.Face
	FontHuakang  font.Face
}

var (
	TitleMenuString = "按回车开始"
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
	FontObject = f
}

//const x = 20
//
//// Draw info
//msg := fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS())
//text.Draw(screen, msg, firaCodeInfo, x, 40, color.White)
//
//// Draw the sample text
//text.Draw(screen, sampleText, firaCodeInfo, x, 80, color.White)
//
//// Draw Kanji text lines
//for i, line := range strings.Split(string(kanjiText), "\n") {
//text.Draw(screen, line, huakang, x, 160+54*i, kanjiTextColor)
//}
