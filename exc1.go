package main

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	var NameImg string = os.Args[1]

	img := gocv.IMRead(NameImg, 1)
	defer img.Close()

	if img.Empty() {
		fmt.Println("img vazia")
	}

	window := gocv.NewWindow("sample")
	defer window.Close()

	window.IMShow(img)
	window.WaitKey(0)
}
