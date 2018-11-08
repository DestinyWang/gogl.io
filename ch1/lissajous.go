package main

import (
	"image/color"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0	// 画板中的第一种颜色
	blackIndex = 1 	// 画板中的第二种颜色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		//handler := func() {}
	}
}

