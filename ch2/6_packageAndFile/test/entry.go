package main

import (
	"fmt"
	"study/ch2/6_packageAndFile/tempconv"
)

func main() {

	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}