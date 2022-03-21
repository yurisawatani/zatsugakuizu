package sushi

import (
	"fmt"
)

type Sushi struct {
	Number  string
	Choices string
	Size    int
	Weight  float64
}

func (a *Sushi) Taberareru() string {
	return fmt.Sprintf("%v%v\n\n", a.Number, a.Choices)
}

func (b *Sushi) mawattekuru() string {
	return fmt.Sprintf("%v流れていった\n %v", b.Number, b.Choices)
}

func main() {
	x := Sushi{
		Number:  "わさび",
		Choices: "まぐろ",
		Size:    8,
		Weight:  50.0,
	}

	y := Sushi{
		Number:  "ネギ",
		Choices: "かつお",
		Size:    7,
		Weight:  40.5,
	}

	x.Taberareru()
	y.mawattekuru()
}
