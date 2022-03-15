package main

import (
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yurisawatani/zatsugakuizu/sushi"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	tate int = 5
	yoko int = 8
)

var (
	mPlus1Regular_ttf font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	ft, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	mPlus1Regular_ttf = ft
}

type Game struct {
	Nyannchudanyan sushi.Sushi
	Wannwann       sushi.Sushi
	Msg            string
	count          int
	Witch          bool
	keys           []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.count = g.count + 1
	if g.count < 60 {
		return nil
	}
	g.count = 0
	if g.Witch {
		g.Witch = false
		g.Msg = g.Nyannchudanyan.Taberareru()
	} else {
		g.Witch = true
		g.Msg = g.Wannwann.Taberareru()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	text.Draw(screen, g.Msg, mPlus1Regular_ttf, 60, 120, color.White)
	for i, k := range g.keys {
		posY := i + 90
		posX := i + 192
		ka := k.String()
		blue := color.RGBA{
			R: 60,
			G: 150,
			B: 255,
			A: 255,
		}
		text.Draw(screen, ka, mPlus1Regular_ttf, posY, posX, blue)
	}
	text.Draw(screen, "3+5=", mPlus1Regular_ttf, 0, 24, color.White)
	if len(g.keys) > 0 {
		answer := g.keys[0].String()
		s := answer
		if strings.HasPrefix(s, "Digit") {
			s = answer[5:]
		}
		text.Draw(screen, s, mPlus1Regular_ttf, 70, 24, color.White)
		if answer == "Digit8" {
			red := color.RGBA{
				R: 255,
				G: 100,
				B: 60,
				A: 255,
			}
			text.Draw(screen, "正解！", mPlus1Regular_ttf, 70, 55, red)
		} else {
			blue2 := color.RGBA{
				R: 60,
				G: 150,
				B: 255,
				A: 255,
			}
			text.Draw(screen, "残念！", mPlus1Regular_ttf, 70, 55, blue2)
		}
		text.Draw(screen, "FINAL ANAWER?", mPlus1Regular_ttf, 150, 230, color.White)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}

func main() {
	ebiten.SetWindowSize(yoko*100, tate*100)
	ebiten.SetWindowTitle("雑学")
	game := &Game{
		Nyannchudanyan: sushi.Sushi{
			Number:  "1.",
			Choices: "まぐろ",
			Size:    8,
			Weight:  50.0,
		},
		Wannwann: sushi.Sushi{
			Number:  "2.",
			Choices: "かつお",
			Size:    7,
			Weight:  40.5,
		},
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
