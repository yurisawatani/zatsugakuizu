package sushi

import (
	"fmt"
)

type Sushi struct {
	Yakumi string
	Fish   string
	Size   int
	Weight float64
}

func (a *Sushi) Taberareru() string {
	return fmt.Sprintf("%vと共に食べられた\n %v", a.Yakumi, a.Fish)
}

func (b *Sushi) mawattekuru() string {
	return fmt.Sprintf("%v流れていった\n %v", b.Yakumi, b.Fish)
}

func main() {
	x := Sushi{
		Yakumi: "わさび",
		Fish:   "まぐろの寿司",
		Size:   8,
		Weight: 50.0,
	}

	y := Sushi{
		Yakumi: "ネギ",
		Fish:   "かつおの寿司",
		Size:   7,
		Weight: 40.5,
	}

	x.Taberareru()
	y.mawattekuru()
}
