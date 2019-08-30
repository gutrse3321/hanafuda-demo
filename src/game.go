package src

import (
	"fmt"
	"font/resources/fonts"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

var (
	sampleText     = `The quick brown fox jumps over the lazy dog.`
	firaCodeInfo   font.Face
	huakang        font.Face
	counter        = 0
	kanjiText      []rune
	kanjiTextColor color.RGBA
)

var jaKanjis []rune

type Game struct {
}

func (g *Game) init() {
}

func init() {
	// table is the list of Japanese Kanji characters in a part of JIS X 0208.
	const table = `
去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房
去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房
去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房
去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房
去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房
去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房
去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房去玩儿浮点数汉族女晶体管费处女座限价房
`
	for _, c := range table {
		if c == '\n' {
			continue
		}
		jaKanjis = append(jaKanjis, c)
	}
}

func init() {
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
	firaCodeInfo = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	huakang = truetype.NewFace(ttH, &truetype.Options{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (g *Game) Update(screen *ebiten.Image) error {
	// Change the text color for each second.
	if counter%ebiten.MaxTPS() == 0 {
		kanjiText = []rune{}
		for j := 0; j < 4; j++ {
			for i := 0; i < 8; i++ {
				kanjiText = append(kanjiText, jaKanjis[rand.Intn(len(jaKanjis))])
			}
			kanjiText = append(kanjiText, '\n')
		}

		kanjiTextColor.R = 0x80 + uint8(rand.Intn(0x7f))
		kanjiTextColor.G = 0x80 + uint8(rand.Intn(0x7f))
		kanjiTextColor.B = 0x80 + uint8(rand.Intn(0x7f))
		kanjiTextColor.A = 0xff
	}
	counter++

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	const x = 20

	// Draw info
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS())
	text.Draw(screen, msg, firaCodeInfo, x, 40, color.White)

	// Draw the sample text
	text.Draw(screen, sampleText, firaCodeInfo, x, 80, color.White)

	// Draw Kanji text lines
	for i, line := range strings.Split(string(kanjiText), "\n") {
		text.Draw(screen, line, huakang, x, 160+54*i, kanjiTextColor)
	}
	return nil
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}
