package main

import (
	"encoding/json"
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yurisawatani/zatsugakuizu/qstrage"
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

type QAP struct {
	Question string
	Answer   string
}

type Game struct {
	Msg             string
	keys            []ebiten.Key
	QuestionlistC   []QAP
	QuestionnumberC uint
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	if len(g.QuestionlistC) == 0 {
		return UpdateStage(g)
	} else {
		return UpdateQuestion(g)
	}
}

func UpdateStage(g *Game) error {
	if len(g.keys) > 0 {
		akey := g.keys[0]
		s := strings.TrimPrefix(akey.String(), "Digit")
		if s == "1" {
			data, err := qstrage.ReadJson("cmd/upload/monndai.json")
			if err != nil {
				log.Fatal(err)
			}
			var listC []QAP
			json.Unmarshal([]byte(data), &listC)
			if err := json.Unmarshal([]byte(data), &listC); err != nil {
				log.Fatal(err)
			}
			g.QuestionlistC = listC
		}
		if s == "2" {
			data, err := qstrage.ReadJson("cmd/upload/monndai.json")
			if err != nil {
				log.Fatal(err)
			}
			var listC []QAP
			json.Unmarshal([]byte(data), &listC)
			if err := json.Unmarshal([]byte(data), &listC); err != nil {
				log.Fatal(err)
			}
			g.QuestionlistC = listC
		}
		return nil
	}
	return nil
}

func UpdateQuestion(g *Game) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	k := len(g.QuestionlistC)
	if k > 0 {
		DrawQuestion(g, screen)
		return
	}
	text.Draw(screen, "ざつがくいず!!\n\nstage 1 --1\n\nstage 2 --2", mPlus1Regular_ttf, 10, 50, color.White)
}

func DrawQuestion(g *Game, screen *ebiten.Image) {
	k := len(g.QuestionlistC)
	if g.QuestionnumberC == uint(k) {
		yelow := color.RGBA{
			R: 255,
			G: 255,
			B: 60,
			A: 255,
		}
		text.Draw(screen, "全問正解！！", mPlus1Regular_ttf, 130, 100, yelow)
		return
	}
	t := g.QuestionlistC[g.QuestionnumberC]
	q := t.Question
	a := t.Answer
	text.Draw(screen, q, mPlus1Regular_ttf, 0, 24, color.White)
	if len(g.keys) > 0 {
		akey := g.keys[0]
		s := strings.TrimPrefix(akey.String(), "Digit")
		text.Draw(screen, s, mPlus1Regular_ttf, 5, 180, color.White)
		if s == a {
			g.QuestionnumberC = g.QuestionnumberC + 1
		} else {
			blue2 := color.RGBA{
				R: 60,
				G: 150,
				B: 255,
				A: 255,
			}
			text.Draw(screen, "残念！", mPlus1Regular_ttf, 70, 180, blue2)
		}
		text.Draw(screen, "FINAL ANSWER?", mPlus1Regular_ttf, 150, 230, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}

func main() {
	data, err := qstrage.ReadJson("cmd/upload/monndai.json")
	if err != nil {
		log.Fatal(err)
	}
	var listC []QAP
	json.Unmarshal([]byte(data), &listC)
	if err := json.Unmarshal([]byte(data), &listC); err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(800, 500)
	ebiten.SetWindowTitle("クイズ")
	gameC := &Game{}
	if err := ebiten.RunGame(gameC); err != nil {
		log.Fatal(err)
	}
}
