package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/yurisawatani/zatsugakuizu/qstrage"
)

func main() {
	filepath := os.Args[1]
	jstr, err := qstrage.ReadJson(filepath)
	if err != nil {
		log.Fatalf("%v", err)
	}
	ctx := context.Background()
	if err := qstrage.WhiteFile(ctx, filepath, strings.NewReader(jstr)); err != nil {
		log.Fatalf("%v", err)
	}
}
