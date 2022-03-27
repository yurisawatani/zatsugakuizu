package main

import (
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yurisawatani/zatsugakuizu/choco"
	"github.com/yurisawatani/zatsugakuizu/hyoushi"
	"github.com/yurisawatani/zatsugakuizu/oishii"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
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
	Msg             string
	count           int
	keys            []ebiten.Key
	QuestionlistC   []choco.Choco
	QuestionnumberC uint
	QuestionlistO   []oishii.Oishii
	QuestionnumberO uint
	QuestionlistH   []hyoushi.Hyoushi
	QuestionnumberH uint
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.count = g.count + 1
	if g.count < 180 {
		return nil
	}
	g.count = 0
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	k := len(g.QuestionlistH)
	if g.QuestionnumberH == uint(k) {
		yelow := color.RGBA{
			R: 255,
			G: 255,
			B: 60,
			A: 255,
		}
		text.Draw(screen, "全問正解！！", mPlus1Regular_ttf, 130, 100, yelow)
		return
	}
	t := g.QuestionlistH[g.QuestionnumberH]
	q := t.Question
	a := t.Answer
	text.Draw(screen, q, mPlus1Regular_ttf, 0, 24, color.White)
	if len(g.keys) > 0 {
		akey := g.keys[0]
		s := strings.TrimPrefix(akey.String(), "Digit")
		text.Draw(screen, s, mPlus1Regular_ttf, 70, 100, color.White)
		if s == a {
			g.QuestionnumberH = g.QuestionnumberH + 1
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
			text.Draw(screen, "残念！", mPlus1Regular_ttf, 70, 130, blue2)
		}
		text.Draw(screen, "FINAL ANSWER?", mPlus1Regular_ttf, 150, 230, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}

func main() {
	ebiten.SetWindowSize(800, 500)
	ebiten.SetWindowTitle("クイズ")
	game := &Game{
		QuestionlistH: hyoushi.Xlist,
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
