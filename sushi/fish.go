package sushi

import (
	"fmt"
)

type Sushi struct {
	Fish   string
	Size   int
	Weight float64
}

func (a *Sushi) Taberareru() {
	fmt.Println("食べられた:", a.Fish)
}

func (b *Sushi) mawattekuru() {
	fmt.Println("流れていった:", b.Fish)
}

func main() {
	x := Sushi{
		Fish:   "まぐろ",
		Size:   8,
		Weight: 50.0,
	}

	y := Sushi{
		Fish:   "はまち",
		Size:   7,
		Weight: 40.5,
	}

	x.Taberareru()
	y.mawattekuru()
}
